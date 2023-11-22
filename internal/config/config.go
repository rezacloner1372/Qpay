package config

import (
	"Qpay/internal/db"
	"Qpay/internal/server"
	"Qpay/pkg/jwt"
)

type Config struct {
	Server *server.Config `json:"server"`
	DB     *db.Config     `koanf:"database"`
	JWT    *jwt.Config    `koanf:"jwt"`
}
