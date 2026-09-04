package handler

import (
	"time"

	"github.com/Narvdeshwar/AetherPay/services/auth/internal/repository"
	"github.com/gin-gonic/gin"
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
	
}
func (h *AuthHandler) Login(c *gin.Context) {

}
