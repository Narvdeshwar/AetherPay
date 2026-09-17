package config

import (
	"fmt"
	"log"

	"github.com/Narvdeshwar/AetherPay/services/payment/internal/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg *Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", cfg.Postgres.Host, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Name, cfg.Postgres.Port, cfg.Postgres.SSLMode, cfg.Postgres.Timezone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("❌ Database connection failed: %v", err)
	}
	// Create dedicated payment schema
	db.Exec("CREATE SCHEMA IF NOT EXISTS payment_schema;")

	// AutoMigrate payments table
	if err := db.AutoMigrate(&repository.Payment{}); err != nil {
		log.Fatalf("❌ AutoMigrate failed: %v", err)
	}

	log.Println("🐘 PostgreSQL Connected & Payment Schema Ready!")
	return db

}
