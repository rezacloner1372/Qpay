package cmd

import (
	"Qpay/internal/config"
	"fmt"
	"os"

	"github.com/labstack/echo"
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
	e := echo.New()

	address := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	fmt.Println(address)
	if err := e.Start(address); err != nil {
		e.Logger.Info("shutting down the server")
	}
}
