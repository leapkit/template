package config

import (
	"os"

	"github.com/joho/godotenv"
)

// envOr returns the value of an environment variable if
// it exists, otherwise it returns the default value
func envOr(name, def string) string {
	// Load .env file only once
	loadOnce.Do(func() {
		godotenv.Load(".env")
	})

	if value := os.Getenv(name); value != "" {
		return value
	}

	return def
}
