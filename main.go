package main

import (
	"Qpay/internal/db"
	"fmt"
)

func init() {
	// Initialize database connection
	dbInstance, err := db.New()
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	// Auto migrate tables
	err = db.AutoMigrate(dbInstance)
	if err != nil {
		fmt.Println("Error during auto migration:", err)
		return
	}

	defer dbInstance.Close()
}

func main() {

}
