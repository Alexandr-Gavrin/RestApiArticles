package api

import "ServerAndDB2/storage"

// General instance for API server

type Config struct {
	// Port
	BindAddr string `toml:"bind_addr"`
	// Logger Level
	LoggerLevel string `toml:"logger_level"`
	// Store config
	Storage *storage.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8080",
		LoggerLevel: "debug",
		Storage:     storage.NewConfig(),
	}
}
