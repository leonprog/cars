package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DbConfig DbConfig
}

func New() *Config {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	dbConfig := newDb()

	return &Config{
		DbConfig: dbConfig,
	}
}

func getEnv(key, defaultValue string) string {
	value, existes := os.LookupEnv(key)
	if !existes {
		return defaultValue
	}
	return value
}
