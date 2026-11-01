package config

import (
	"fmt"
	"os"
)

type Config struct {
	HTTPAddr    string
	DatabaseURL string
}

func Load() (Config, error) {
	return Config{
		HTTPAddr:    getenv("HTTPAddr", ":8080"),
		DatabaseURL: createDatabaseURL(),
	}, nil
}

func getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}

func createDatabaseURL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		getenv("POSTGRES_USER", "admin"),
		getenv("POSTGRES_PASSWORD", "admin"),
		getenv("DB_HOST", "localhost"),
		getenv("DB_PORT", "5432"),
		getenv("POSTGRES_DB", "scope"),
		getenv("SSL_MODE", "disable"),
	)
}
