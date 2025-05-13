package config

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func LoadEnv(envFilename string) error {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Join(filepath.Dir(filename), "..")
	envPath := filepath.Join(dir, envFilename)
	err := godotenv.Load(envPath)
	if err != nil {
		return err
	}
	return nil
}

func NewConfig() *Config {
	return &Config{
		Key: os.Getenv("APIKEY"),
	}
}

var Setup = NewConfig()
