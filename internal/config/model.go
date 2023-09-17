package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func New() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("errorr when load env: %s", err.Error())
	}

	return &Config{
		Server{
			Host: os.Getenv("SRV_HOST"),
			Port: os.Getenv("SRV_PORT"),
		},
		Database{
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASS"),
			Name: os.Getenv("DB_NAME"),
		},
	}
}
