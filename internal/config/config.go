package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	Env        string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	JWTSecret string
}

func Load() Config {

	_ = godotenv.Load()

	cfg := Config{
		ServerPort: getEnv("SERVER_PORT", "8080"),
		Env:        getEnv("APP_ENV", "development"),

		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", ""),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", ""),

		JWTSecret: getEnv("JWT_SECRET", "defaultsecret"),
	}

	log.Println("Settings loaded successfully!")
	return cfg
}


func getEnv(key string, defaultValue string) string {
	val, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return val
}
