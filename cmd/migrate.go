package cmd

import (
	"Qpay/internal/config"
	"Qpay/internal/db"
	"Qpay/internal/model"
	"Qpay/internal/repository"
	"Qpay/pkg/logger"
	"os"

	"github.com/sirupsen/logrus"
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
	logger := logger.NewLogger(cfg.Logger)

	if len(args) != 1 {
		logger.WithFields(logrus.Fields{
			"args": args,
		}).Fatal("Invalid arguments given")
	}

	db, err := db.New()
	if err != nil {
		logger.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("Error creating db")
	}

	repo := repository.New(logger, db)
	if err := repo.Migrate(model.Migrate(args[0])); err != nil {
		logger.WithFields(logrus.Fields{
			"migration": map[string]interface{}{
				"arg":   args[0],
				"error": err.Error(),
			},
		}).Fatal("Error migrating")
	}

	logger.WithFields(logrus.Fields{
		"migration": args[0],
	}).Info("Database has been migrated successfully")
}
