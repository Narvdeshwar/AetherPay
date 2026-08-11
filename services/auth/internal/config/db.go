package config

import (
	"log"

	"github.com/AetherPay/services/auth/internal/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgrespassword dbname=aetherpay port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect with database:%v", err)
	}
	db.Exec("CREATE SCHEMA IF NOT EXISTS auth_schema;")
	err = db.AutoMigrate(&repository.Merchant{})
	if err != nil {
		log.Fatalf("Error performing the migration %v", err)
	}
	log.Println("🐘 PostgreSQL Connected & Schema Migrated Successfully!")
	return db
}
