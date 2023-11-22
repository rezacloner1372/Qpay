package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
	dsn := "root:123456@tcp(qpay-mysql:3306)/qpay?charset=utf8mb4"
	// Open a Gorm DB connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
