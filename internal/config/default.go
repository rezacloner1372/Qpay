package config

import (
	"Qpay/internal/db"
	"Qpay/internal/server"
	"Qpay/pkg/logger"
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
		Logger: &logger.Config{
			Development: true,
			Encoding:    "json",
			Level:       "info",
		},
		// JWT: &jwt.Config{
		// 	Secret:     "be8f3d8d-d01b-4d76-987f-42d219de9f6b",
		// 	Expiration: 24,
		// },
	}
}
