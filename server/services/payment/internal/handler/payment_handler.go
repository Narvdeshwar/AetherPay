package handler

import (
	"github.com/Narvdeshwar/AetherPay/services/payment/internal/repository"
	"github.com/redis/go-redis/v9"
)

type PaymentHandler struct {
	repo repository.NewPaymentRepository
	rdb  *redis.Client
}

func NewPaymentHandler(repo repository.PaymentRepository) *PaymentHandler {
	return &PaymentHandler{
		repo: repo,
	}
}
