package database

import (
	"log"
	"os"

	"github.com/kevinanielsen/opinionated/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	dsn := os.Getenv("DSN")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Poll{})
	db.AutoMigrate(models.Question{})
	db.AutoMigrate(models.Vote{})

	DB = db
	log.Println("Successfully connected to database!")
}
