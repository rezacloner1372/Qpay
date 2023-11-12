package config

import "Qpay/internal/db"

type Config struct {
	DB *db.Config `koanf:"database"`
}
