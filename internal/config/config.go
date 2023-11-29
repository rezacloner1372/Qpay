package config

import (
	"Qpay/internal/db"
	"Qpay/internal/server"
	"Qpay/pkg/bank"
	"Qpay/pkg/logger"
)

type Config struct {
	Server *server.Config `koanf:"server"`
	DB     *db.Config     `koanf:"database"`
	Logger *logger.Config `koanf:"logger"`
	// JWT    *jwt.Config    `koanf:"jwt"`
	Bank *bank.Config `koanf:"bank"`
}
