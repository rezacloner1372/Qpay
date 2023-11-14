package config

import (
	"Qpay/internal/db"
	"Qpay/internal/server"
)

func Default() *Config {
	return &Config{
		Server: &server.Config{
			Host: "localhost",
			Port: 8080,
		},
		DB: &db.Config{
			Host:     "localhost",
			Port:     3306,
			Username: "root",
			Password: "123456",
			Database: "MYSQL",
		},
	}
}
