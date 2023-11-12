package config

import (
	"Qpay/internal/db"
)

func Default() *Config {
	return &Config{
		DB: &db.Config{
			Host:     "localhost",
			Port:     3306,
			Username: "root",
			Password: "123456",
			Database: "MYSQL",
		},
	}
}
