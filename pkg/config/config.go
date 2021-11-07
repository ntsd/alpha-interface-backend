package config

import (
	"log"
	"os"
	"strings"
)

// Config is struct for application configuration
type Config struct {
	Environment string
	Port        string
	DevMode     bool
}

// GetConfig function to configuration struct
func GetConfig() Config {
	config := Config{
		Environment: requireEnv("ENVIRONMENT"),
		Port:        defaultENV("PORT", "8001"),
		DevMode:     strings.ToLower(os.Getenv("DEV_MODE")) == "true",
	}

	return config
}

// requireEnv use to get required env. This will fatal when empty
func requireEnv(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		log.Fatalf("Required env %s not found", key)
	}
	return value
}

// defaultENV get env if none return default
func defaultENV(key, df string) string {
	value := os.Getenv(key)
	if value == "" {
		return df
	}
	return value
}
