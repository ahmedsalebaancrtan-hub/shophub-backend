package infra

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port        string
	DBPort      string
	DBUser      string
	DBHost      string
	DBName      string
	DBPassword  string
}

var Configuration AppConfig

func InitEnv() error {

	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("failed to load environment variables: %w", err)
	}

	Configuration = AppConfig{
		Port:       os.Getenv("PORT"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBHost:     os.Getenv("DB_HOST"),
		DBName:     os.Getenv("DB_NAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
	}

	// Validate required environment variables
	if Configuration.Port == "" {
		return fmt.Errorf("PORT is required")
	}

	if Configuration.DBHost == "" ||
		Configuration.DBPort == "" ||
		Configuration.DBUser == "" ||
		Configuration.DBName == "" ||
		Configuration.DBPassword == "" {
		return fmt.Errorf("database environment variables are required")
	}

	return nil
}