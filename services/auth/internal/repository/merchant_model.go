package repository

import "time"

// Merchant represents the auth_schema.merchants table
type Merchant struct {
	ID           string    `gorm:"primaryKey;type:varchar(64)" json:"id"`
	TenantId     string    `gorm:"uniqueIndex;type:varchar(64);not null" json:"tenant_id"`
	MerchantName string    `gorm:"type:varchar(255);not null" json:"merchant_name"`
	Email        string    `gorm:"uniqueIndex;type:varchar(255);not null" json:"email"`
	PasswordHash string    `gorm:"type:varchar(255);not null" json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// TableName explicitly defines schema-per-service naming
func (Merchant) TableName() string {
	return "auth_schema.merchants"
}
