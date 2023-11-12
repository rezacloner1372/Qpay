package cmd

import (
	"Qpay/internal/config"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Migrate struct{}

func (m Migrate) Command(trap chan os.Signal) *cobra.Command {
	run := func(_ *cobra.Command, args []string) {
		m.main(config.Load(), args, trap)
	}

	return &cobra.Command{
		Use:       "migrate",
		Short:     "run database migrations",
		Run:       run,
		Args:      cobra.OnlyValidArgs,
		ValidArgs: []string{"up", "down"},
	}
}

func (m Migrate) main(config *config.Config, args []string, trap chan os.Signal) {
	//Todo: we run Migration function in here
	fmt.Println("Hello Migarte")
}
