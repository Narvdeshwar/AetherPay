package config

import (
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	// postgresql
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBSSLMode  string
	DBTimezone string

	//JWT
	JWTSecret        string
	AuthPort         string
	JWTExpiryMinutes time.Duration

	// REDIS
	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int

	// Rate limit time
	PublicRateLimit     int
	PublicRateLimitTime time.Duration

	ProtectedRateLimit     int
	ProtectedRateLimitTime time.Duration
}

func LoadConfig() *Config {
	_ = godotenv.Load(".env")

	return &Config{
		DBHost:     getEnv("HOST_ADDRESS"),
		DBUser:     getEnv("POSTGRES_USER"),
		DBPassword: getEnv("POSTGRES_PASSWORD"),
		DBName:     getEnv("POSTGRES_DB"),
		DBPort:     getEnv("POSTGRES_PORT"),
		DBSSLMode:  getEnv("DB_SSLMODE"),
		DBTimezone: getEnv("DB_TIMEZONE"),

		// JWT
		JWTSecret:        getEnv("JWT_SECRET"),
		AuthPort:         getEnv("AUTH_PORT"),
		JWTExpiryMinutes: getDuration("JWT_EXPIRY"),

		// Redis
		RedisHost:     getEnv("REDIS_HOST"),
		RedisPort:     getEnv("REDIS_PORT"),
		RedisPassword: getEnv("REDIS_PASSWORD"),
		RedisDB:       getInt("REDIS_DB"),

		// Rate Limiting
		PublicRateLimit:     getInt("PUBLIC_RATE_LIMIT"),
		PublicRateLimitTime: getDuration("PUBLIC_RATE_LIMIT_WINDOW"),

		ProtectedRateLimit:     getInt("PROTECTED_RATE_LIMIT"),
		ProtectedRateLimitTime: getDuration("PROTECTED_RATE_LIMIT_WINDOW"),
	}
}
