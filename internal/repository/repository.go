package repository

import (
	"Qpay/internal/model"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository interface {
	Migrate(model.Migrate) error

	//Todo: Implement another database CRUD functions
}

type repository struct {
	logger *logrus.Logger
	db     *gorm.DB
}

func New(logger *logrus.Logger, db *gorm.DB) Repository {
	r := &repository{logger: logger, db: db}

	return r
}

func (r repository) Migrate(direction model.Migrate) error {
	instace, _ := r.db.DB()
	driver, err := mysql.WithInstance(instace, &mysql.Config{})
	if err != nil {
		return err
	}

	mig, err := migrate.NewWithDatabaseInstance(
		"file://./internal/db/migration",
		"qpay",
		driver)
	if err != nil {
		return err
	}

	if direction == model.Up {
		if err := mig.Up(); err != nil {
			return err
		}
	} else {
		if err := mig.Down(); err != nil {
			return err
		}
	}

	return nil
}
