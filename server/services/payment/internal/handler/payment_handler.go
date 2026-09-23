package handler

import (
	"github.com/Narvdeshwar/AetherPay/services/payment/internal/repository"
	"github.com/redis/go-redis/v9"
)

type PaymentHandler struct {
	repo repository.PaymentRepository
	rdb  *redis.Client
}

func NewPaymentHandler(repo repository.PaymentRepository, rdb *redis.Client) *PaymentHandler {
	return &PaymentHandler{
		repo: repo,
		rdb:  rdb,
	}
}

type CreatePaymentRequest struct {
	Amount   int64  `json:"amount" binding:"required,gt=0"`
	Currency string `json:"currency", binding:"required"`
}
