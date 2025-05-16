package database

import (
	"fmt"

	"github.com/janction/meme-launchpad/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(dsn string) error {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	return nil
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&models.Token{})
}
