package repository

import "gorm.io/gorm"

type MerchantRepository interface {
	Create(merchant *Merchant) error
}

type merchantRepository struct {
	db *gorm.DB
}

func NewMerchantRepository(db *gorm.DB) MerchantRepository {
	return &merchantRepository{db: db}
}

func (r *merchantRepository) Create(m *Merchant) error {
	return r.db.Create(m).Error
}
