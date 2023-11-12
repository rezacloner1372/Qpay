package config

import "Qpay/internal/db"

type Config struct {
	//Todo
	DB db.Config `koanf:"database"`
}
