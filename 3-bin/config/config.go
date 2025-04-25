package config

import (
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func NewConfig() *Config {
	godotenv.Load(".env")
	return &Config{
		Key: os.Getenv("KEY"),
	}
}
