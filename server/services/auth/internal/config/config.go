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
		DBHost:     getEnv("HOST_ADDRESS", "db_key"),
		DBUser:     getEnv("POSTGRES_USER", "db_user"),
		DBPassword: getEnv("POSTGRES_PASSWORD", "postgre_pass"),
		DBName:     getEnv("POSTGRES_DB", "db_name"),
		DBPort:     getEnv("POSTGRES_PORT", "db_port"),
		DBSSLMode:  getEnv("DB_SSLMODE", "db_ssl_mode"),
		DBTimezone: getEnv("DB_TIMEZONE", "db_time_zone"),

		// JWT
		JWTSecret:        getEnv("JWT_SECRET", "jwt_secret"),
		AuthPort:         getEnv("AUTH_PORT", "auth_port"),
		JWTExpiryMinutes: getDuration("JWT_EXPIRY_MINUTES", "jwt_expiry"),

		// Redis
		RedisHost:     getEnv("REDIS_HOST", "redis_host"),
		RedisPort:     getEnv("REDIS_PORT", "redis_port"),
		RedisPassword: getEnv("REDIS_PASSWORD", "redis_password"),
		RedisDB:       getInt("REDIS_DB", "redis_db"),

		// Rate Limiting
		PublicRateLimit:     getInt("PUBLIC_RATE_LIMIT", "public_rate_limit"),
		PublicRateLimitTime: getDuration("PUBLIC_RATE_LIMIT_WINDOW", "public_rate_limit_window"),

		ProtectedRateLimit:     getInt("PROTECTED_RATE_LIMIT", "protected_rate_limit"),
		ProtectedRateLimitTime: getDuration("PROTECTED_RATE_LIMIT_WINDOW", "protected_rate_limit_window"),
	}
}
