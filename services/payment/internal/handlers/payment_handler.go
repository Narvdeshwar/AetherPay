package handlers

import (
	"aetherpay/payment/internal/repository"
	"aetherpay/shared"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	Repo *repository.PaymentRepository
}

func NewPaymentHandler(repo *repository.PaymentRepository) *PaymentHandler {
	return &PaymentHandler{Repo: repo}
}

func (h *PaymentHandler) ProcessPayment(c *gin.Context) {
	var req shared.PaymentRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		
	}
}
