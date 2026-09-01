package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomeClaims struct {
	UserID   string `json:"user_id"`
	TenantID string `json:"tenant_id"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(userID, tenantID, role, secret string, duration time.Duration) (string, error) {
	claims := CustomeClaims{
		UserID:   userID,
		TenantID: tenantID,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "aetherpay-auth-service",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func main() {
	message, err := GenerateToken("rtest", "test", "admin", "usr", 2)
	if err != nil {
		fmt.Println("Error encountered")
	}
	fmt.Println("Token Generated:", message)
}
