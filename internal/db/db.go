package db

import (
	"Qpay/internal/model"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	dbConn *gorm.DB
)

func CreateDBConnection() (*gorm.DB, error) {

	if dbConn != nil {
		CloseDBConnection(dbConn)
	}

	dsn := "root:123456@tcp(mysql:3306)/qpay?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open("mysql", dsn)

	if err != nil {
		fmt.Println("Failed to connect database")
		return nil, err
	}

	db.DB().SetConnMaxIdleTime(time.Minute * 5)
	db.DB().SetMaxIdleConns(10)

	db.LogMode(true)
	fmt.Println("Connected database")

	dbConn = db
	return db, nil
}

func GetDatabaseConnection() *gorm.DB {
	sqlDB := dbConn.DB()

	if sqlDB == nil {
		fmt.Println("Failed to get database connection")
	}

	if err := sqlDB.Ping(); err != nil {
		fmt.Println("Failed to ping database connection")
	}

	return dbConn
}

func CloseDBConnection(db *gorm.DB) {
	sqlDB := db.DB()

	if sqlDB == nil {
		fmt.Println("Failed to close database connection")
	}

	defer sqlDB.Close()
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
