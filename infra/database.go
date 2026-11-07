package infra

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() error {

	config := Configuration

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s dbname=%s sslmode=disable",
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBPort,
		config.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	DB = db

	return nil
}