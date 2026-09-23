package main

import (
	"log"

	"github.com/Narvdeshwar/AetherPay/services/payment/internal/config"
	"github.com/Narvdeshwar/AetherPay/services/payment/internal/handler"
	"github.com/Narvdeshwar/AetherPay/services/payment/internal/repository"
	"github.com/Narvdeshwar/AetherPay/shared/database"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Env files are not loaded successfully %v", err.Error())
	}
	db := config.InitDB(cfg)
	paymentRepo := repository.NewPaymentRepository(db)
	rdb, err := database.InitRedis(cfg.Redis)
	if err != nil {
		log.Fatalf("failed to initialize Redis: %v", err)
	}

	paymentHandler := handler.NewPaymentHandler(paymentRepo, rdb)
	r := gin.Default()

	log.Println("auth Service is running on port 8002")
	v1 := r.Group("/api/v1/payments")
	{
		v1.POST("/process-payment", paymentHandler.ProcessPayment)
	}
	if err := r.Run(cfg.PaymentPort); err != nil {
		log.Fatalf("Error running the auth server %v", err)
	}
}
