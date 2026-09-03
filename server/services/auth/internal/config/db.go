package config

import (
	"fmt"
	"log"

	"github.com/Narvdeshwar/AetherPay/services/auth/internal/repository"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg *Config) *gorm.DB {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSSLMode, cfg.DBTimezone)
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("X database connection failed! %v", err)
	}
	// Schema & Table Auto-Creation
	err = db.AutoMigrate(&repository.Merchant{})
	if err != nil {
		log.Fatalf("Database migration failed! %v", err)
	}
	log.Println("*****************PostgreSQL Connected & Merchants Table Ready!************")
	return db
}
