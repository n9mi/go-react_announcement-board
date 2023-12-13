package database

import (
	"fmt"
	"server/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDBConnString(config config.DBConfig) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort,
	)
}

func NewDB(config config.DBConfig) (*gorm.DB, error) {
	dsn := getDBConnString(config)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}
