package main

import (
	"Qpay/cmd"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

func main() {
	const description = "QPay application"
	root := &cobra.Command{Short: description}

	trap := make(chan os.Signal, 1)
	signal.Notify(trap, syscall.SIGINT, syscall.SIGTERM)

	root.AddCommand(
		cmd.Migrate{}.Command(trap),
	)

	if err := root.Execute(); err != nil {
		log.Fatalf("failed to execute root command: \n%v", err)
	}
}
