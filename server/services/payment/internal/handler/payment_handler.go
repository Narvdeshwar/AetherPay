// services/payment/internal/handler/payment_handler.go
package handler

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Narvdeshwar/AetherPay/services/payment/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type PaymentHandler struct {
	repo repository.PaymentRepository
	rdb  *redis.Client
}

func NewPaymentHandler(repo repository.PaymentRepository, rdb *redis.Client) *PaymentHandler {
	return &PaymentHandler{repo: repo, rdb: rdb}
}

type CreatePaymentRequest struct {
	Amount   int64  `json:"amount" binding:"required,gt=0"` // Must be greater than 0
	Currency string `json:"currency" binding:"required"`
}

func (h *PaymentHandler) ProcessPayment(c *gin.Context) {

	ctx := c.Request.Context()

	// --------------------------------------------------
	// 1. Idempotency key
	// --------------------------------------------------

	idempotencyKey := strings.TrimSpace(
		c.GetHeader("Idempotency-Key"),
	)

	if idempotencyKey == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing 'Idempotency-Key' in headers",
		})
		return
	}

	// --------------------------------------------------
	// 2. Tenant
	// --------------------------------------------------

	tenantID := c.GetHeader("X-Tenant-ID")

	if tenantID == "" {
		tenantID = "tn_merchant_demo"
	}

	// TODO:
	// Production:
	// tenantID should come from JWT middleware,
	// NOT from X-Tenant-ID.

	redisKey := fmt.Sprintf(
		"idempotency:%s:%s",
		tenantID,
		idempotencyKey,
	)

	// --------------------------------------------------
	// 3. Fast Redis lock
	// --------------------------------------------------

	locked, err := h.rdb.SetNX(
		ctx,
		redisKey,
		"PROCESSING",
		2*time.Minute,
	).Result()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Redis lock check failed",
		})
		return
	}

	// --------------------------------------------------
	// 4. Existing request
	// --------------------------------------------------

	if !locked {

		existingPayment, err :=
			h.repo.FindByIdempotencyKey(
				idempotencyKey,
				tenantID,
			)

		if err == nil {

			c.JSON(http.StatusOK, gin.H{
				"message": "Payment already processed",
				"payment": existingPayment,
			})

			return
		}

		// Redis says another request is processing.
		val, err := h.rdb.Get(ctx, redisKey).Result()

		if err == nil && val == "PROCESSING" {

			c.JSON(http.StatusConflict, gin.H{
				"error":   "Duplicate request",
				"message": "Payment is currently processing. Please wait.",
			})

			return
		}
	}

	// --------------------------------------------------
	// 5. Parse request
	// --------------------------------------------------

	var req CreatePaymentRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		h.rdb.Del(ctx, redisKey)

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})

		return
	}

	req.Currency = strings.ToUpper(
		strings.TrimSpace(req.Currency),
	)

	// --------------------------------------------------
	// 6. Check DB AGAIN
	// --------------------------------------------------

	existingPayment, err :=
		h.repo.FindByIdempotencyKey(
			idempotencyKey,
			tenantID,
		)

	if err == nil {

		h.rdb.Set(
			ctx,
			redisKey,
			"SUCCESS",
			24*time.Hour,
		)

		c.JSON(http.StatusOK, gin.H{
			"message": "Payment already processed",
			"payment": existingPayment,
		})

		return
	}

	// --------------------------------------------------
	// 7. Mock payment gateway
	// --------------------------------------------------

	time.Sleep(1 * time.Second)

	// --------------------------------------------------
	// 8. Create payment
	// --------------------------------------------------

	payment := &repository.Payment{
		ID:             "pay_" + uuid.NewString(),
		TenantID:       tenantID,
		IdempotencyKey: idempotencyKey,
		Amount:         req.Amount,
		Currency:       req.Currency,
		Status:         "SUCCESS",
		CreatedAt:      time.Now(),
	}

	// --------------------------------------------------
	// 9. PostgreSQL is final protection
	// --------------------------------------------------

	if err := h.repo.Create(payment); err != nil {

		// Important:
		// If duplicate idempotency key,
		// fetch existing payment and return it.

		existingPayment, findErr :=
			h.repo.FindByIdempotencyKey(
				idempotencyKey,
				tenantID,
			)

		if findErr == nil {

			h.rdb.Set(
				ctx,
				redisKey,
				"SUCCESS",
				24*time.Hour,
			)

			c.JSON(http.StatusOK, gin.H{
				"message": "Payment already processed",
				"payment": existingPayment,
			})

			return
		}

		h.rdb.Del(ctx, redisKey)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to record payment",
		})

		return
	}

	// --------------------------------------------------
	// 10. Redis SUCCESS
	// --------------------------------------------------

	h.rdb.Set(
		ctx,
		redisKey,
		"SUCCESS",
		24*time.Hour,
	)

	// --------------------------------------------------
	// 11. Response
	// --------------------------------------------------

	c.JSON(http.StatusCreated, gin.H{
		"message": "Payment successful",
		"payment": payment,
	})
}
