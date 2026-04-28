package repository

import (
	"aetherpay/shared"

	"gorm.io/gorm"
)

type PaymentRepository struct {
	DB *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{
		DB: db,
	}
}

func (r *PaymentRepository) SaveTransaction(txn *shared.Transaction) error {
	result := r.DB.Create(txn)
	return result.Error
}
