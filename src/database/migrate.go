package database

import (
	"log"
	"os"

	"github.com/kevinanielsen/opinionated/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Migrate() {
	dsn := os.Getenv("DSN")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(models.User{}, models.Poll{}, models.Question{}, models.Vote{})

	log.Println("Successfully migrated database!")
}
