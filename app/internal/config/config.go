package config

import (
	"os"
)

// Config holds application configuration
type Config struct {
	Addr             string
	BaseURL          string // e.g. http://localhost:8080
	DatabaseUser     string
	DatabasePassword string
	DatabaseHost     string
	DatabasePort     string
	DatabaseName     string
}

// Load reads configuration from environment with sensible defaults
func Load() Config {
	cfg := Config{
		Addr:             getEnv("ADDR", ":8080"),
		BaseURL:          getEnv("BASE_URL", "http://localhost:8080"),
		DatabaseUser:     getEnv("DB_USER", "user"),
		DatabasePassword: getEnv("DB_PASS", "pass"),
		DatabaseHost:     getEnv("DB_HOST", "localhost"),
		DatabasePort:     getEnv("DB_PORT", "5432"),
		DatabaseName:     getEnv("DB_NAME", "dbname"),
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
