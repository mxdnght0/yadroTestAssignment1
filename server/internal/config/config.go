package config

import (
	"os"
)

type Config struct {
	path string
	port string
}

func NewConfig() *Config {
	path := os.Getenv("RESOLV_PATH")
	if path == "" {
		path = "/etc/resolv.conf"
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8000"
	}

	return &Config{
		path: path,
		port: port,
	}
}

func (c *Config) GetPath() string {
	return c.path
}

func (c *Config) GetPort() string {
	return c.port
}
