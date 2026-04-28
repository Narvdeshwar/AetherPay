package handlers

import (
	"aetherpay/shared"
	"net/http"

	"github.com/gin-gonic/gin"
    // Apne path ke hisab se module name daalein
	"billing/internal/service"
)

type BillingHandler struct {
	Service *service.BillingService
}

func NewBillingHandler(svc *service.BillingService) *BillingHandler {
	return &BillingHandler{Service: svc}
}

func (h *BillingHandler) Subscribe(c *gin.Context) {
	var req shared.PaymentRequest
	
	// Ab request user (Postman) se aayegi, hardcoded nahi!
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Service (Brain) ko call karein
	resp, err := h.Service.SubscribeUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Subscription failed", "details": err.Error()})
		return
	}

	// Response Bhejein
	if resp.Status == "SUCCESS" {
		c.JSON(http.StatusOK, gin.H{
			"message":        "subscription successful",
			"transaction_id": resp.TransactionId,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":        "payment failed",
			"transaction_id": resp.TransactionId,
		})
	}
}