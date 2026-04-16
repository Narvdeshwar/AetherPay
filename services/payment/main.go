package main

import (
	"aetherpay/shared"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.POST("/api/v1/payment/process", func(c *gin.Context) {
		var req shared.PaymentRequest
		if err := c.ShouldBindBodyWithJSON(&req); err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}
		fmt.Printf("Processing the payment request of %s %.2f by the User: %s (%s)", req.Currency, req.Amount, req.Name, req.UserId)
		resp := shared.PaymentResponse{
			TransactionId: "1x23fdfg",
			Status:        "SUCCESS",
		}
		c.JSON(http.StatusOK, resp)
	})

	fmt.Println("Server is running on the port 3002")
	app.Run(":3002")
}
