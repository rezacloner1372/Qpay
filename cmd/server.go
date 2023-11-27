package cmd

import (
	"Qpay/internal/config"
	"Qpay/internal/server"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

type Server struct{}

func (s Server) Command(trap chan os.Signal) *cobra.Command {
	cfg, _ := config.Load(false)
	run := func(_ *cobra.Command, _ []string) {
		s.main(cfg, trap)
	}

	return &cobra.Command{
		Use:   "server",
		Short: "run server",
		Run:   run,
	}
}

func (s *Server) main(cfg *config.Config, trap chan os.Signal) {
	address := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	server.NewServer().Start(address)

	// Keep this at the bottom of the main function
	// Wait for a signal
	sigReceived := <-trap

	// Build the log message
	message := "Exiting by receiving a Unix signal"
	field := "Signal Trap: " + sigReceived.String()

	// Log the message with the field
	log.Printf("%s [%s]\n", message, field)
}
