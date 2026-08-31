package main

import (
	"log"

	"github.com/AetherPay/services/auth/internal/config"
	"github.com/AetherPay/services/auth/internal/handler"
	"github.com/AetherPay/services/auth/internal/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.InitDB()
	merchantRepo = repository.NewMerchantRepository(db)
	authHandler := handler.NewAuthHandler(merchantRepo)
	
	r := gin.Default()
	v1 := r.Group("/api/v1/auth")
	{
		v1.POST("/login", authHandler.Login)
	}
	log.Println("Auth service is running on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error running the server restart the server again or check for error:%v", err)
	}
}
