package config

import (
	"time"

	sharedconfig "github.com/Narvdeshwar/AetherPay/shared/config"
)

type Config struct {
	// Common configurations
	Postgres  sharedconfig.PostgresConfig
	Redis     sharedconfig.RedisConfig
	RateLimit sharedconfig.RateLimitConfig

	// Auth-specific configurations
	JWTSecret string
	AuthPort  string
	JWTExpiry time.Duration
}

func LoadConfig() (*Config, error) {
	_ = sharedconfig.LoadDotEnv()

	redisDB, err := sharedconfig.GetInt(
		"REDIS_DB",
		0,
	)
	if err != nil {
		return nil, err
	}

	jwtExpiry, err := sharedconfig.GetDuration(
		"JWT_EXPIRY",
		15*time.Minute,
	)
	if err != nil {
		return nil, err
	}

	publicLimit, err := sharedconfig.GetInt(
		"PUBLIC_RATE_LIMIT",
		5,
	)
	if err != nil {
		return nil, err
	}

	publicWindow, err := sharedconfig.GetDuration(
		"PUBLIC_RATE_LIMIT_WINDOW",
		time.Minute,
	)
	if err != nil {
		return nil, err
	}

	protectedLimit, err := sharedconfig.GetInt(
		"PROTECTED_RATE_LIMIT",
		10,
	)
	if err != nil {
		return nil, err
	}

	protectedWindow, err := sharedconfig.GetDuration(
		"PROTECTED_RATE_LIMIT_WINDOW",
		time.Minute,
	)
	if err != nil {
		return nil, err
	}

	return &Config{
		Postgres: sharedconfig.PostgresConfig{
			Host:     sharedconfig.GetEnv("DB_HOST", "localhost"),
			User:     sharedconfig.GetEnv("DB_USER", "postgres"),
			Password: sharedconfig.GetEnv("DB_PASSWORD", ""),
			Name:     sharedconfig.GetEnv("DB_NAME", "aetherpay"),
			Port:     sharedconfig.GetEnv("DB_PORT", "5433"),
			SSLMode:  sharedconfig.GetEnv("DB_SSLMODE", "disable"),
			Timezone: sharedconfig.GetEnv("DB_TIMEZONE", "UTC"),
		},

		Redis: sharedconfig.RedisConfig{
			Host:     sharedconfig.GetEnv("REDIS_HOST", "localhost"),
			Port:     sharedconfig.GetEnv("REDIS_PORT", "6379"),
			Password: sharedconfig.GetEnv("REDIS_PASSWORD", ""),
			DB:       redisDB,
		},

		RateLimit: sharedconfig.RateLimitConfig{
			PublicLimit:     publicLimit,
			PublicWindow:    publicWindow,
			ProtectedLimit:  protectedLimit,
			ProtectedWindow: protectedWindow,
		},

		JWTSecret: sharedconfig.GetEnv("JWT_SECRET", ""),
		AuthPort:  sharedconfig.GetEnv("AUTH_PORT", "3001"),
		JWTExpiry: jwtExpiry,
	}, nil
}
