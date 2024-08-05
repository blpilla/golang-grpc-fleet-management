package config

import (
	"log"
	"os"
)

// Config struct to hold configuration values
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

// LoadConfig function to load configuration from environment variables
func LoadConfig() *Config {
	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "fleet_management"),
	}
}

// Helper function to get environment variables with a fallback value
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// Initialize function to log the loaded configuration
func Initialize() {
	config := LoadConfig()
	log.Printf("Database Host: %s", config.DBHost)
	log.Printf("Database Port: %s", config.DBPort)
	log.Printf("Database User: %s", config.DBUser)
	log.Printf("Database Name: %s", config.DBName)
}
