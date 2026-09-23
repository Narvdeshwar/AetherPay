package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Narvdeshwar/AetherPay/services/payment/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type PaymentHandler struct {
	repo repository.PaymentRepository
	rdb  *redis.Client
}

func NewPaymentHandler(repo repository.PaymentRepository, rdb *redis.Client) *PaymentHandler {
	return &PaymentHandler{
		repo: repo,
		rdb:  rdb,
	}
}

type CreatePaymentRequest struct {
	Amount   int64  `json:"amount" binding:"required,gt=0"`
	Currency string `json:"currency" binding:"required"`
}

func (h *PaymentHandler) ProcessPayment(c *gin.Context) {
	ctx := context.Background()
	// 1. Mandatory Idempotency-Key Header
	idempotencyKey := c.GetHeader("Idempotency-Key")
	if idempotencyKey == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'Idempotency-Key' in headers"})
		return
	}
	// 2. Tenant Context (In production, extracted from JWT middleware)
	tenantID := c.GetHeader("X-Tenant-ID")
	if tenantID == "" {
		tenantID = "tn_merchant_demo"
	}
	redisKey := fmt.Sprintf("idempotency:%s:%s", tenantID, idempotencyKey)
	// 3. Redis Lock (SETNX - Set if Not Exists)
	// Agar key pehle se nahi hai toh 'PROCESSING' lock lagao (TTL 2 minutes)
	locked, err := h.rdb.SetNX(ctx, redisKey, "PROCESSING", 2*time.Minute).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Redis lock check failed"})
		return
	}
	// Agar key pehle se exist karti hai:
	if !locked {
		val, _ := h.rdb.Get(ctx, redisKey).Result()
		if val == "PROCESSING" {
			c.JSON(http.StatusConflict, gin.H{
				"error":   "Duplicate request",
				"message": "Payment is currently processing. Please wait.",
			})
			return
		}
		// Agar already SUCCESS tha, DB se purana payment return karo (No double charge!)
		existingPayment, err := h.repo.FindByIdempotencyKey(idempotencyKey)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "Payment already processed (Idempotent replay)",
				"payment": existingPayment,
			})
			return
		}
	}

	// 4. Parse Request
	var req CreatePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.rdb.Del(ctx, redisKey) // Invalid payload par lock release karo
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	// 5. Payment Simulation (Mock Gateway Transaction)
	time.Sleep(1 * time.Second) // Simulating bank API call delay

	payment := &repository.Payment{
		ID:             fmt.Sprintf("pay_%d", time.Now().UnixNano()%1000000),
		TenantID:       tenantID,
		IdempotencyKey: idempotencyKey,
		Amount:         req.Amount,
		Currency:       req.Currency,
		Status:         "SUCCESS",
		CreatedAt:      time.Now(),
	}

	if err := h.repo.Create(payment); err != nil {
		h.rdb.Del(ctx, redisKey)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to record payment"})
		return
	}

	// 6. Redis update: Mark as SUCCESS with 24-hour retention
	h.rdb.Set(ctx, redisKey, "SUCCESS", 24*time.Hour)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Payment successful!",
		"payment": payment,
	})

}
