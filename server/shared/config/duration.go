package config

import (
	"fmt"
	"os"
	"time"
)

// GetDuration reads a duration environment variable.
//
// Supported values:
// 15m
// 1h
// 30s
// 500ms
func GetDuration(
	key string,
	defaultValue time.Duration,
) (time.Duration, error) {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue, nil
	}

	parsedValue, err := time.ParseDuration(value)
	if err != nil {
		return 0, fmt.Errorf(
			"invalid duration value for %s: %w",
			key,
			err,
		)
	}

	return parsedValue, nil
}
