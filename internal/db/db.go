package db

import (
	"Qpay/internal/model"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func New() (*gorm.DB, error) {
	dsn := "root:123456@tcp(mysql:3306)/qpay?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Failed to connect database")
		return nil, err
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
	fmt.Println("Connected database")
	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&model.User{},
		&model.Transactions{},
		&model.Tariffs{},
		&model.Roles{},
		&model.RolesPermissions{},
		&model.Permissions{},
		&model.PaymentGateways{},
	).Error

	if err != nil {
		fmt.Println("Failed to migrate database")
		return err
	}

	fmt.Println("Migrated database")
	return nil
}
