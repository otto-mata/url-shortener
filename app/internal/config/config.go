package config

import (
	"os"
)

// Config holds application configuration
type Config struct {
	Addr       string
	BaseURL    string // e.g. http://localhost:8080
	ValkeyHost string
	ValkeyPort string
}

// Load reads configuration from environment with sensible defaults
func Load() Config {
	cfg := Config{
		Addr:       getEnv("ADDR", ":8080"),
		BaseURL:    getEnv("BASE_URL", "http://localhost:8080"),
		ValkeyHost: getEnv("VALKEY_HOST", "localhost"),
		ValkeyPort: getEnv("VALKEY_PORT", "6379"),
	}
	return cfg
}

func getEnv(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}
