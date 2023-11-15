package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// create a new instance of the logger based on the provided configuration.
func NewLogger(cfg *Config) *logrus.Logger {
	logger := logrus.New()

	// Set log level
	level, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logger.SetLevel(level)

	// Set log formatter based on encoding
	if cfg.Encoding == "json" {
		logger.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logger.SetFormatter(&logrus.TextFormatter{})
	}

	// Set other options, e.g., report caller
	logger.SetReportCaller(true)

	// Set log output
	logger.SetOutput(os.Stdout)

	return logger
}
