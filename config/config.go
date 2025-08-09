package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type NatsConfig struct {
	Host string
	Port string
}

type Config struct {
	Env  string
	Db   DbConfig
	Nats NatsConfig
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file. Loading default values")
	}

	return &Config{
		Env: getEnv("ENV", "dev"),
		Db: DbConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Name:     getEnv("DB_NAME", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
		},
		Nats: NatsConfig{
			Host: getEnv("NATS_HOST", "localhost"),
			Port: getEnv("NATS_PORT", "4222"),
		},
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
