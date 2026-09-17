package main

import (
	"log"

	"github.com/Narvdeshwar/AetherPay/services/payment/internal/config"
	"github.com/Narvdeshwar/AetherPay/services/payment/internal/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Env files are not loaded successfully %v", err.Error())
	}
	db := config.InitDB(cfg)
	paymentRepo := repository.NewPaymentRepository(db)
	payHandler := handler.NewPaymentHandler(paymentRepo)
	r := gin.Default()

	log.Println("auth Service is running on port 3001")
	if err := r.Run(cfg.PaymentPort); err != nil {
		log.Fatalf("Error running the auth server %v", err)
	}
}
