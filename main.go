package main

import (
	"Qpay/cmd"
	"Qpay/pkg/logger"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	const description = "QPay application"
	root := &cobra.Command{Short: description}

	trap := make(chan os.Signal, 1)
	signal.Notify(trap, syscall.SIGINT, syscall.SIGTERM)

	root.AddCommand(
		cmd.Server{}.Command(trap),
		cmd.Migrate{}.Command(trap),
	)

	if err := root.Execute(); err != nil {
		log.Fatalf("failed to execute root command: \n%v", err)
	}

	cfg := logger.DefaultConfig()
	qpaylogger := logger.NewLogger(cfg)

	// sample usage.
	qpaylogger.Info("This is an informational message", logrus.Fields{"key": "value"})
	qpaylogger.Warn("This is a warning message", logrus.Fields{"pi": 3.14159})
	qpaylogger.Error("This is an error message", logrus.Fields{"error": "something went wrong"})
	qpaylogger.Debug("This is a debug message", logrus.Fields{"debugInfo": "some debug info"})

}
