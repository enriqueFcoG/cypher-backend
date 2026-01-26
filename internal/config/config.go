package config

import "os"

type Config struct {
	Env      string
	Port     string
	Database string
}

func Load() Config {
	cfg := Config{
		Env:      os.Getenv("APP_ENV"),
		Port:     os.Getenv("APP_PORT"),
		Database: os.Getenv("APP_DATABASE"),
	}
	return cfg
}
