package config

import (
	"log"
	"os"
	"strconv"
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
}

func LoadConfig() *Config {
	_ = godotenv.Load(".env")
	// important to parse the time into the time.ParseDuration since we are expecting the time in the minutes
	jwtExpiryMinutes, err := time.ParseDuration(os.Getenv("JWT_EXPIRY_MINUTES"))
	if err != nil {
		log.Fatalf("jwtsecret is not a number:%v", err)
	}
	redisDBStr := os.Getenv("REDIS_DB")
	redisDB, err := strconv.Atoi(redisDBStr)
	if err != nil {
		log.Fatalf("REDIS_DB is not a valid number: %v", err)
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
		RedisHost:        os.Getenv("REDIS_HOST"),
		RedisPort:        os.Getenv("REDIS_PORT"),
		RedisPassword:    os.Getenv("REDIS_PASSWORD"),
		RedisDB:          redisDB,
	}
}
