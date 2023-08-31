package config

import "os"

// Config.
type Config struct {
	HTTPAddr    string
	DatabaseUrl string
}

// Read reads config from environment.
func Read() Config {
	var config Config

	httpAddr, exists := os.LookupEnv("HTTP_ADDR")
	if exists {
		config.HTTPAddr = httpAddr
	}
	databaseUrl, exists := os.LookupEnv("DATABASE_URL")
	if exists {
		config.DatabaseUrl = databaseUrl
	}

	return config
}
