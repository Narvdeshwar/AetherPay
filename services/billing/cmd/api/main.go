package main

import (
	"aetherpay/billing/internal/clients"
	"aetherpay/billing/internal/handlers"
	service "aetherpay/billing/internal/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Initialize HTTP Client
	// Docker mein 'payment-service' container name use hoga
	paymentClient := clients.NewPaymentClient("http://payment-service:3002")

	// 2. Initialize Service & Handler
	svc := service.NewBillingService(paymentClient)
	handler := handlers.NewBillingHandler(svc)

	// 3. Setup Router
	app := gin.Default()
	api := app.Group("/api/v1")
	{
		api.POST("/billing/subscribe", handler.Subscribe)
	}

	// 4. Start Server
	fmt.Println("🧾 Billing Service running on port 3001")
	app.Run(":3001")
}