package main

import (
	"log"

	"github.com/AetherPay/services/auth/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	authHandler := handler.NewAuthHandler()
	v1 := r.Group("/api/v1/auth")
	{
		v1.POST("/login", authHandler.Login)
	}
	log.Println("Auth service is running on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error running the server restart the server again or check for error:%v", err)
	}
}
