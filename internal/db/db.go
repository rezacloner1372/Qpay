package db

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbConn *gorm.DB
)

func CreateDBConnection(cfg *Config) (*gorm.DB, error) {
	if dbConn != nil {
		return dbConn, nil // Reuse existing connection
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	// Open a Gorm DB connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetConnMaxIdleTime(time.Minute * 5)
	sqlDB.SetMaxIdleConns(10)

	dbConn = db
	return dbConn, nil
}

func GetDatabaseConnection() (*gorm.DB, error) {
	if dbConn != nil {
		sqlDB, err := dbConn.DB()
		if err != nil {
			return nil, err
		}

		if err := sqlDB.Ping(); err != nil {
			return nil, err
		}

		return dbConn, nil
	}

	return nil, fmt.Errorf("failed to get database connection")
}

func CloseDBConnection() error {
	if dbConn != nil {
		sqlDB, err := dbConn.DB()
		if err != nil {
			return err
		}

		err = sqlDB.Close()
		if err != nil {
			return err
		}

		dbConn = nil
	}

	return nil
}
