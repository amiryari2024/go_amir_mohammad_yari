package config

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// InitDB initializes the database connection
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return nil, err
	}
	fmt.Println("Database connection established")
	return db, nil
}
