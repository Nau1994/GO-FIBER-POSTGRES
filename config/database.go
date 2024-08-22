package config

import (
	"GoFiber/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	dsn := "host=localhost user=postgres password=root dbname=app port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// AutoMigrate User and Post models
	DB.AutoMigrate(&models.User{}, &models.Post{})
	log.Println("Database connection established and migrated")
}
