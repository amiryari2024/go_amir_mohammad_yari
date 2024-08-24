package config

import (
	"myproject/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("myproject.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}
	DB.AutoMigrate(&models.User{}, &models.Package{})
}
