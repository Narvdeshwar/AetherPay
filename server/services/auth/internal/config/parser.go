package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("%s key is not set", value)
	}
	return value
}

func getInt(key string) int {
	value := getEnv(key)
	result, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("%s must be value integer value:%v", key, err)
	}
	return result

}

func getDuration(key string) time.Duration {
	value := getEnv(key)
	result, err := time.ParseDuration(value)
	if err != nil {
		log.Fatalf("%s must be valid duration:%v", key, err)
	}
	return result
}
