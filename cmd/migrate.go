package cmd

import (
	"Qpay/internal/config"
	"os"

	"github.com/spf13/cobra"
)

type Migrate struct{}

func (m Migrate) Command(trap chan os.Signal) *cobra.Command {
	cfg, _ := config.Load(true)
	run := func(_ *cobra.Command, args []string) {
		m.main(cfg, args, trap)
	}

	return &cobra.Command{
		Use:       "migrate",
		Short:     "run migrations",
		Run:       run,
		Args:      cobra.OnlyValidArgs,
		ValidArgs: []string{"up", "down"},
	}
}

func (m *Migrate) main(cfg *config.Config, args []string, trap chan os.Signal) {
	//Todo : we run Migration functions here
}
