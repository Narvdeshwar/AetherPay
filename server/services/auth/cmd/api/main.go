package main

import (
	"log"
	"os"

	"github.com/Narvdeshwar/AetherPay/services/auth/internal/config"
	"github.com/Narvdeshwar/AetherPay/services/auth/internal/handler"
	"github.com/Narvdeshwar/AetherPay/services/auth/internal/repository"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.LoadConfig()
	db := config.InitDB(cfg)

	merchantRepo := repository.NewMerchantRepository(db)

	authHandler := handler.NewAuthHandler(merchantRepo, os.Getenv("JWT_SECRET"), 15)

	r := gin.Default()
	v1 := r.Group("/api/v1/auth")
	{
		v1.POST("/register", authHandler.Register)
		v1.POST("/login", authHandler.Login)
	}
	log.Println("auth Service is running on port 3001")
	if err := r.Run(os.Getenv("AUTH_PORT")); err != nil {
		log.Fatalf("Error running the auth server %v", err)
	}

}
