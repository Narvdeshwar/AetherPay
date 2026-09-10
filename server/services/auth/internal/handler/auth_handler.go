package handler

import (
	"net/http"
	"time"

	requestdto "github.com/Narvdeshwar/AetherPay/services/auth/internal/dto/request"
	responsedto "github.com/Narvdeshwar/AetherPay/services/auth/internal/dto/response"
	"github.com/Narvdeshwar/AetherPay/services/auth/internal/repository"
	"github.com/Narvdeshwar/AetherPay/shared"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	repo      repository.MerchantRepository
	jwtSecret string
	tokenTTL  time.Duration
}

// constructor
func NewAuthHandler(repo repository.MerchantRepository, jwtSecret string, tokenTTL time.Duration) *AuthHandler {
	return &AuthHandler{
		repo:      repo,
		jwtSecret: jwtSecret,
		tokenTTL:  tokenTTL,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req requestdto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate the hashed password", "details": err.Error()})
		return
	}
	merchantID := "mrc_" + uuid.NewString()
	tenantID := "tn_" + uuid.NewString()
	merchant := repository.Merchant{
		ID:           merchantID,
		TenantID:     tenantID,
		MerchantName: req.MerchantName,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
	}
	if err := h.repo.Create(&merchant); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Merchant with this email address is already registered", "message": err.Error()})
		return
	}
	res := responsedto.RegisterResponse{
		MerchantResponse: responsedto.MerchantResponse{
			MerchantID:   merchant.ID,
			TenantID:     merchant.TenantID,
			MerchantName: merchant.MerchantName,
			Email:        merchant.Email,
			CreatedAt:    merchant.CreatedAt,
		},
	}

	c.JSON(http.StatusCreated, res)
}
func (h *AuthHandler) Login(c *gin.Context) {
	var req requestdto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	// searching user in the merchant
	merchant, err := h.repo.FindByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(merchant.PasswordHash), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	// token generation and adding the tenant_id
	token, err := shared.GenerateToken(merchant.ID, merchant.TenantID, "admin", h.jwtSecret, h.tokenTTL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in generating the token", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"access_token": token, "tenant_id": merchant.TenantID, "expiry": int64(h.tokenTTL.Seconds())})
}

func (h *AuthHandler) Profile(c *gin.Context) {
	// Go Concept: Type Assertion
	// c.Get() generic 'any' return karta hai, string me convert karne ke liye .(string) use hota hai
	tenantID, exits := c.Get("tenant_id")
	if !exits {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Tenant context missing"})
		return
	}
	userID, _ := c.Get("user_id")
	merchant, err := h.repo.FindByIdAndTenant(userID.(string), tenantID.(string))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user or tenant"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":       "Access granted to protected resource!",
		"tenant_id":     tenantID,
		"user_id":       userID,
		"email":         merchant.Email,
		"merchant_name": merchant.MerchantName,
		"created_at":    merchant.CreatedAt,
		"updated_at":    merchant.UpdatedAt,
	})
}

//
