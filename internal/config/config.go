package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	MySQLHost     string
	MySQLPort     int
	MySQLUsername string
	MySQLPassword string
	MySQLDatabase string
}

func LoadConfig() (*AppConfig, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("MYSQL_HOST")
	portStr := os.Getenv("MYSQL_PORT")
	username := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")
	database := os.Getenv("MYSQL_DATABASE")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("failed to convert MYSQL_PORT to integer: %v", err)
	}

	config := &AppConfig{
		MySQLHost:     host,
		MySQLPort:     port,
		MySQLUsername: username,
		MySQLPassword: password,
		MySQLDatabase: database,
	}

	return config, nil
}
