package config

import (
	"Qpay/internal/db"
	"Qpay/internal/server"
	"Qpay/pkg/utils"
)

type Config struct {
	Server *server.Config `json:"server"`
	DB     *db.Config     `koanf:"database"`
	JWT    *utils.Config  `koanf:"jwt"`
}
