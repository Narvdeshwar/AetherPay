package handlers

import (
	service "aetherpay/payment/internal/services"
	"aetherpay/shared"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	Service *service.PaymentService
}

func NewPaymentHandler(svc *service.PaymentService) *PaymentHandler {
	return &PaymentHandler{Service: svc}
}

func (h *PaymentHandler) ProcessPayment(c *gin.Context) {
	var req shared.PaymentRequest
	
	// 1. Input Check
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. Call The Brain (Service)
	txnId, err := h.Service.ExecutePayment(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process payment"})
		return
	}

	// 3. Success Response
	c.JSON(http.StatusOK, shared.PaymentResponse{TransactionId: txnId, Status: "SUCCESS"})
}