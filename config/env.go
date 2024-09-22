package config

import (
	"os"
)

// Define the Config struct to store application configurations
type Config struct {
	PublicHost string
	Port       string
	DBURL      string // This will store the full PostgreSQL connection URL (DSN)
}

var Envs = initConfig()

// Helper function to get environment variables with default values
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// Initialize the configuration by reading environment variables or using defaults
func initConfig() Config {
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DBURL:      getEnv("DATABASE_URL", "postgresql://postgres:@localhost:5432/ecom?sslmode=disable"),
	}
}
