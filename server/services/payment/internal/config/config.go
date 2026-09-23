package config

import (
	sharedconfig "github.com/Narvdeshwar/AetherPay/shared/config"
)

type Config struct {
	Postgres    sharedconfig.PostgresConfig
	Redis       sharedconfig.RedisConfig
	PaymentPort string
}

func LoadConfig() (*Config, error) {
	_ = sharedconfig.LoadDotEnv()

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
			Host: sharedconfig.GetEnv("REDIS_HOST", "localhost"),
			Port: sharedconfig.GetEnv("REDIS_PORT", "6379"),
		},
		PaymentPort: sharedconfig.GetEnv("PAYMENT_PORT", "8002"),
	}, nil
}
