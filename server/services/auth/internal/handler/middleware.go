package handler

import (
	"net/http"
	"strings"

	"github.com/Narvdeshwar/AetherPay/shared"
	"github.com/gin-gonic/gin"
)

func (h *AuthHandler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Request Header se 'Authorization' nikalo
		authHeader := c.GetHeader("Authorization")
		// Token absent hai toh yahi abort karo (Aage badhne nahi dena)
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missin"})
			return
		}
		// 2. Format check: "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format, expected 'Bearer <token>'"})
			return
		}
		tokenString := parts[1]
		// 3. Shared package se offline cryptographic validation
		claims, err := shared.ValidateJWT(tokenString, h.jwtSecret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token", "details": err.Error()})
			return
		}
		// 4. Go/Gin Concept: Context Storage (Downstream handlers ke liye save karna)
		c.Set("tenant_id", claims.TenantID)
		c.Set("user_id", claims.UserID)

		// 5. Aage jaane ki permission do
		c.Next()
	}
}
