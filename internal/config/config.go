package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port        string
}

func Load() (*Config, error) {
	// load .env from project root
	var err error = godotenv.Load()
	if err != nil {
		return nil, err
	}

	var config *Config = &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        os.Getenv("PORT"),
	}

	return config, nil
}