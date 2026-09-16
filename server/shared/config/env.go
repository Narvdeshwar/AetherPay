package config

import (
	"fmt"
	"os"
	"strconv"
)

// GetEnv returns the environment variable value.
// If the variable is missing, it returns the default value.
func GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}

// GetRequiredEnv returns an error if the environment variable is missing.
func GetRequiredEnv(key string) (string, error) {
	value := os.Getenv(key)

	if value == "" {
		return "", fmt.Errorf(
			"required environment variable %s is missing",
			key,
		)
	}

	return value, nil
}

// GetInt reads an integer environment variable.
func GetInt(key string, defaultValue int) (int, error) {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue, nil
	}

	parsedValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf(
			"invalid integer value for %s: %w",
			key,
			err,
		)
	}

	return parsedValue, nil
}
