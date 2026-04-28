package main

import (
	"aetherpay/payment/internal/broker"
	"aetherpay/payment/internal/config"
	"aetherpay/payment/internal/handlers"
	"aetherpay/payment/internal/repository"
	service "aetherpay/payment/internal/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. DB connection
	config.InitDB()

	// 2. Initialize Components (Wiring)
	repo := repository.NewPaymentRepository(config.DB)
	kafkaPub := broker.NewKafkaPublisher("kafka:29092")
	rabbitPub := broker.NewRabbitMQPublisher("amqp://guest:guest@rabbitmq:5672")

	svc := service.NewPaymentService(repo, kafkaPub, rabbitPub)
	handler := handlers.NewPaymentHandler(svc)

	// 3. Setup router
	app := gin.Default()
	api := app.Group("/api/v1")
	{
		api.POST("/payment/process", handler.ProcessPayment)
	}

	// 4. Start server
	fmt.Println("Payment service running on port 3002")
	app.Run(":3002")
}
