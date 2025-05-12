package config

import (
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func NewConfig() *Config {
	godotenv.Load()
	return &Config{
		Key: os.Getenv("APIKEY"),
	}
}

var Setup = NewConfig()
