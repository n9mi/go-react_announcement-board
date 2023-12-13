package config

import (
	"os"
)

type DBConfig struct {
	DBHost     string
	DBName     string
	DBUser     string
	DBPassword string
	DBPort     string
}

func GetDBConfig() DBConfig {
	return DBConfig{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
	}
}
