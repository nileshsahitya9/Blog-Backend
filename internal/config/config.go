package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DBHost     string
	DBPort     int
	DBUsername string
	DBPassword string
	DBDatabase string
}

func LoadConfig() (*AppConfig, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	portStr := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("failed to convert DB_PORT to integer: %v", err)
	}

	config := &AppConfig{
		DBHost:     host,
		DBPort:     port,
		DBUsername: username,
		DBPassword: password,
		DBDatabase: database,
	}

	return config, nil
}
