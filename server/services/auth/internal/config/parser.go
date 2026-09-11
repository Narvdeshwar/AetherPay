package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

func getEnv(key, key_type string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("%s key is not set of %s", value, key_type)
	}
	return value
}

func getInt(key, key_type string) int {
	value := getEnv(key, key_type)
	result, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("%s must be value integer value:%v", key, err)
	}
	return result

}

func getDuration(key, key_type string) time.Duration {
	value := getEnv(key, key_type)
	result, err := time.ParseDuration(value)
	if err != nil {
		log.Fatalf("%s must be valid duration:%v", key, err)
	}
	return result
}
