package repository

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	ID             string    `gorm:"primaryKey;type:varchar(64)" json:"id"`
	TenantID       string    `gorm:"index;type:varchar(64);not null" json:"tenant_id"`
	IdempotencyKey string    `gorm:"uniqueIndex;type:varchar(128);not null" json:"idempotency_key"`
	Amount         int64     `gorm:"type:bigint;not null" json:"amount"` // In Paisa (₹100 = 10000)
	Currency       string    `gorm:"type:varchar(3);default:'INR'" json:"currency"`
	Status         string    `gorm:"type:varchar(32);not null" json:"status"` // PENDING, SUCCESS, FAILED
	CreatedAt      time.Time `json:"created_at"`
}

func (Payment) TableName() string {
	return "payment_schema.payments"
}

type PaymentRepository interface {
	Create(payment *Payment) error
	FindByIdempotencyKey(key, tenantID string) (*Payment, error)
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db: db}
}

func (r *paymentRepository) Create(p *Payment) error {
	return r.db.Create(p).Error
}

func (r *paymentRepository) FindByIdempotencyKey(
	key string,
	tenantID string,
) (*Payment, error) {

	var payment Payment

	err := r.db.
		Where(
			"idempotency_key = ? AND tenant_id = ?",
			key,
			tenantID,
		).
		First(&payment).Error

	if err != nil {
		return nil, err
	}

	return &payment, nil
}
