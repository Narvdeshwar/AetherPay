package handler

import (
	"net/http"
	"time"

	"github.com/AetherPay/services/auth/internal/repository"
	"github.com/Narvdeshwar/AetherPay/shared"
	"github.com/gin-gonic/gin"
)

const JWT_SECRET_KEY = "super-secret-aether-key-2026"

type AuthHandler struct {
	repo repository.MerchantRepository
}

func NewAuthHandler(repo repository.MerchantRepository) *AuthHandler {
	return &AuthHandler{repo: repo}
}

type RegisterRequest struct{
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	MerchantName string `json:"merchant_name" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// Login Handles merchant auth & issues JWT
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Payload or validation failed", "details": err.Error()})
		return
	}
	// Simulated Merchant Lookup
	mockTenantID := "tn_merchant_991"
	mockUserID := "usr_7721"

	// Token Generation via shared package
	token, err := shared.GenerateJWT(mockUserID, mockTenantID, "admin", JWT_SECRET_KEY, 15*time.Minute)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in generating the token", "err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"access_token": token,
		"tenant_id":    mockTenantID,
		"expires_in":   "15m",
	})
}
