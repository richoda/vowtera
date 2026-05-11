package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv   string
	AppPort  string
	AppName  string
	DBHost   string
	DBPort   string
	DBName   string
	DBUser   string
	DBPass   string
	JWTSecret string
}

func Load() *Config {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	envFile := ".env." + env
	if err := godotenv.Load(envFile); err != nil {
		log.Printf("env file %s tidak ditemukan, pakai env OS", envFile)
	}

	return &Config{
		AppEnv:    getEnv("APP_ENV", "development"),
		AppPort:   getEnv("APP_PORT", "8080"),
		AppName:   getEnv("APP_NAME", "admin-api"),
		DBHost:    getEnv("DB_HOST", "localhost"),
		DBPort:    getEnv("DB_PORT", "5432"),
		DBName:    getEnv("DB_NAME", "vowtera_dev"),
		DBUser:    getEnv("DB_USER", "postgres"),
		DBPass:    getEnv("DB_PASSWORD", ""),
		JWTSecret: getEnv("JWT_SECRET", "changeme"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
