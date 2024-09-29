package config

import (
	"os"
	"strconv"
)

// Define the Config struct to store application configurations
type Config struct {
	PublicHost           string
	Port                 string
	DBURL                string // This will store the full PostgreSQL connection URL (DSN)
	JWTDurationinSeconds int64  // This will
	JWTSecret            string // This will store the secret token
}

var Envs = initConfig()

// Helper function to get environment variables with default values
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}

// Initialize the configuration by reading environment variables or using defaults
func initConfig() Config {
	return Config{
		PublicHost:           getEnv("PUBLIC_HOST", "http://localhost"),
		Port:                 getEnv("PORT", "8080"),
		DBURL:                getEnv("DATABASE_URL", "postgresql://postgres:@localhost:5432/ecom?sslmode=disable"),
		JWTDurationinSeconds: getEnvAsInt("JWT_EXP", 3600*7*24),
		JWTSecret:            getEnv("JWT_SECRET", "this one is secret though"),
	}
}
