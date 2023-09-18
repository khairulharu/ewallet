package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Get() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("errorr when load env: %s", err.Error())
	}

	return &Config{
		Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		Database{
			Host: os.Getenv("DATABASE_HOST"),
			Port: os.Getenv("DATABASE_PORT"),
			User: os.Getenv("DATABASE_USER"),
			Pass: os.Getenv("DATABASE_PASS"),
			Name: os.Getenv("DATABASE_NAME"),
		},
	}
}
