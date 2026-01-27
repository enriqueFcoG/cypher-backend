package config

import (
	"log"
	"os"
)

type Config struct {
	Env      string
	Port     string
	Database string
}

func Load() Config {
	cfg := Config{
		Env:      getEnv("APP_ENV", "dev"),
		Port:     getEnv("APP_PORT", "8081"),
		Database: mustGetEnv("APP_DATABASE"),
	}
	return cfg
}

func getEnv(key string, default_val string) string {
	val := os.Getenv(key)

	if val == "" {
		return default_val
	}
	return val
}

func mustGetEnv(key string) string {
	val := os.Getenv(key)

	if val == "" {
		log.Fatalf("error getting value with key %s", key)
	}

	return val
}
