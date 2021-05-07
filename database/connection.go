package database

import (
	"github.com/danielmoisa/go-jwt/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(sqlite.Open("dev.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	DB = connection
	connection.AutoMigrate(&models.User{})
}
