package config

import (
	"Qpay/internal/db"
	"Qpay/internal/server"
)

type Config struct {
	Server *server.Config `json:"server"`
	DB     *db.Config     `koanf:"database"`
}
