package config

import "time"

type RateLimitConfig struct {
	PublicLimit       int
	PublicWindow      time.Duration
	ProtectedLimit    int
	ProtectedWindow   time.Duration
}
