package main

import (
	"log"

	"github.com/Narvdeshwar/AetherPay/services/auth/internal/config"
	"github.com/Narvdeshwar/AetherPay/services/auth/internal/handler"
	"github.com/Narvdeshwar/AetherPay/services/auth/internal/middleware"

	"github.com/Narvdeshwar/AetherPay/services/auth/internal/repository"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.LoadConfig()
	db := config.InitDB(cfg)

	merchantRepo := repository.NewMerchantRepository(db)

	authHandler := handler.NewAuthHandler(merchantRepo, cfg.JWTSecret, cfg.JWTExpiryMinutes)

	r := gin.Default()
	// public route
	public := r.Group("/api/v1/auth")
	{
		public.POST("/register", authHandler.Register)
		public.POST("/login", authHandler.Login)
	}

	// protected route
	protected := r.Group("/api/v1/auth")
	protected.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	{
		protected.GET("/profile", authHandler.Profile)
	}
	log.Println("auth Service is running on port 3001")
	if err := r.Run(cfg.AuthPort); err != nil {
		log.Fatalf("Error running the auth server %v", err)
	}

}
