package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBSSLMode  string
	DBTimezone string
}

func LoadConfig() *Config {
	_ = godotenv.Load(".env")
	return &Config{
		DBHost:     os.Getenv("HOST_ADDRESS"),
		DBUser:     os.Getenv("POSTGRES_USER"),
		DBPassword: os.Getenv("POSTGRES_PASSWORD"),
		DBName:     os.Getenv("POSTGRES_DB"),
		DBPort:     os.Getenv("POSTGRES_PORT"),
		DBSSLMode:  os.Getenv("DB_SSLMODE"),
		DBTimezone: os.Getenv("DB_TIMEZONE"),
	}
}
