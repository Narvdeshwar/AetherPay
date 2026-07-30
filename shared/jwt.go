package shared

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserId   string `json:"user_id"`
	TenantId string `json:"tenantId"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(userId, tenantId, role, secret string, duration time.Duration) (string, error) {
	claims := CustomClaims{
		UserId:   userId,
		TenantId: tenantId,
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
