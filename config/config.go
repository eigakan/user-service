package config

import (
	"fmt"
	"os"
	"strconv"

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
	Port int
}

type JwtConfig struct {
	Secret   string
	ExpHours int16
}

type Config struct {
	Env  string
	Db   DbConfig
	Nats NatsConfig
	Jwt  JwtConfig
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file. Loading default values")
	}

	tokenExpHours, err := strconv.Atoi(getEnv("JWT_EXP_HOURS", "24"))
	if err != nil {
		panic("Cannot parse JWT_EXP_HOURS from env. Please set it to a valid integer value")
	}

	natsPort, err := strconv.Atoi(getEnv("NATS_PORT", "4222"))
	if err != nil {
		panic("Cannot parse JWT_EXP_HOURS from env. Please set it to a valid integer value")
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
			Port: natsPort,
		},

		Jwt: JwtConfig{
			Secret:   getEnv("JWT_SECRET", "my-super-secret"),
			ExpHours: int16(tokenExpHours),
		},
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
