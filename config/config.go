package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type ServerConfig struct {
	Port string
}

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

var AppConfig *Config

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file : %v", err)
	}

	AppConfig = &Config{
		Server: ServerConfig{
			Port: getEnv("PORT", "3000"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "127.0.0.1"),
			Port:     getEnv("DB_PORT", "3306"),
			Username: getEnv("DB_USER", "root"),
			Password: getEnv("DB_PASSWORD", ""),
			Database: getEnv("DB_NAME", "db_clean_go"),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}
	return defaultValue
}
