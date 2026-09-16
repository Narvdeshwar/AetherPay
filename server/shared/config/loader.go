package config

import (
	"github.com/joho/godotenv"
)

func LoadDotEnv() error {
	return godotenv.Load(".env")
}
