package infra

import (
	"os"

	"github.com/joho/godotenv"
)

// Config holds the application configuration
type Config struct {
	Port              string
	GinMode           string
	DatabaseURL       string
	LineChannelId     string
	LineChannelSecret string
	LogLevel          string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	// Load .env file if it exists
	_ = godotenv.Load()

	config := &Config{
		Port:              getEnv("PORT", "8080"),
		GinMode:           getEnv("GIN_MODE", "debug"),
		DatabaseURL:       getEnv("DATABASE_URL", ""),
		LineChannelId:     getEnv("LINE_CHANNEL_ID", ""),
		LineChannelSecret: getEnv("LINE_CHANNEL_SECRET", ""),
		LogLevel:          getEnv("LOG_LEVEL", "debug"),
	}

	return config, nil
}

// getEnv gets an environment variable with a fallback value
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
