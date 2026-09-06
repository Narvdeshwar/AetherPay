package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost           string
	DBUser           string
	DBPassword       string
	DBName           string
	DBPort           string
	DBSSLMode        string
	DBTimezone       string
	JWTSecret        string
	AuthPort         string
	JWTExpiryMinutes time.Duration
}

func LoadConfig() *Config {
	_ = godotenv.Load(".env")
	// important to parse the time into the time.ParseDuration since we are expecting the time in the minutes
	jwtExpiryMinutes, err := time.ParseDuration(os.Getenv("JWT_EXPIRY_MINUTES"))
	if err != nil {
		log.Fatalf("jwtsecret is not a number:%v", err)
	}
	return &Config{
		DBHost:           os.Getenv("HOST_ADDRESS"),
		DBUser:           os.Getenv("POSTGRES_USER"),
		DBPassword:       os.Getenv("POSTGRES_PASSWORD"),
		DBName:           os.Getenv("POSTGRES_DB"),
		DBPort:           os.Getenv("POSTGRES_PORT"),
		DBSSLMode:        os.Getenv("DB_SSLMODE"),
		DBTimezone:       os.Getenv("DB_TIMEZONE"),
		JWTSecret:        os.Getenv("JWT_SECRET"),
		AuthPort:         os.Getenv("AUTH_PORT"),
		JWTExpiryMinutes: jwtExpiryMinutes,
	}
}
