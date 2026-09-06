package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/Narvdeshwar/AetherPay/services/auth/internal/repository"
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

type RegisterRequest struct {
	Email        string `json:"email" binding:"required,email"`
	Password     string `json:"password" binding:"required,min=8"`
	MerchantName string `json:"merchant_name" binding:"required,min=2,max=255"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	log.Println("Register request received", req)

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
	c.JSON(http.StatusCreated, gin.H{
		"message":     "Merchant Registered Successfully",
		"tenant_id":   merchant.TenantID,
		"merchant_id": merchant.ID,
	})
}
func (h *AuthHandler) Login(c *gin.Context) {

}
