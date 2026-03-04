package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseUrl string
	JWTSecret   string
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func LoadConfig() *Config {
	_ = godotenv.Load()
	return &Config{
		Port:        getEnv("PORT", "8080"),
		DatabaseUrl: getEnv("DATABASE_URL", "postgresql://root:password@localhost:5432/postgres"),
		JWTSecret:   getEnv("JWT_SECRET", "your-jwt-secret"),
	}
}
