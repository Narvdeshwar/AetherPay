package repository

import (
	"time"

	"gorm.io/gorm"
)

// Go Concept: Struct Mapping to SQL Table
type Merchant struct {
	ID           string    `gorm:"primaryKey;type:varchar(64)" json:"id"`
	TenantID     string    `gorm:"uniqueIndex;type:varchar(64);not null" json:"tenant_id"`
	MerchantName string    `gorm:"type:varchar(255);not null" json:"merchant_name"`
	Email        string    `gorm:"uniqueIndex;type:varchar(255);not null" json:"email"`
	PasswordHash string    `gorm:"type:varchar(255);not null" json:"-"` // "-" means JSON response me password kabhi nahi dikhega
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// Go Concept: Interfaces (Polymorphism & Abstraction)
// Interface batata hai ki kya kaam karna hai, kaise karna hai wo Struct decide karega.
type MerchantRepository interface {
	Create(m *Merchant) error
	FindByEmail(email string) (*Merchant, error)
}

type merchantRepository struct {
	db *gorm.DB // Go Concept: Pointer to DB Connection pool
}

// Go Concept: Constructor Pattern (New...)
func NewMerchantRepository(db *gorm.DB) MerchantRepository {
	return &merchantRepository{db: db}
}

func (r *merchantRepository) Create(m *Merchant) error {
	return r.db.Create(m).Error
}

func (r *merchantRepository) FindByEmail(email string) (*Merchant, error) {
	var m Merchant
	err := r.db.Where("email = ?", email).First(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}
