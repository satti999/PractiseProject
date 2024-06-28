package config

import (
	"os"
)

type Configuration struct {
	Port string
}

var Config Configuration

func LoadConfig() {
	Config = Configuration{
		Port: getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		value = fallback
	}
	return value
}
