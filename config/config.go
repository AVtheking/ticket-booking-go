package config

import (
	"fmt"
	"os"
)

type Config struct {
	Database DBConfig
}

type DBConfig struct {
	PostgresConnectionString string
}

func NewConfig() *Config {
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")

	return &Config{
		Database: DBConfig{
			PostgresConnectionString: fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME),
		},
	}
}
