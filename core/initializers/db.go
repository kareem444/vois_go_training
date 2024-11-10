package initializers

import (
	"fmt"
	"time"

	"example.com/test/core/env"
	"example.com/test/core/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error

	DB_HOST := env.Get("DB_HOST")
	DB_PORT := env.Get("DB_PORT")
	DB_USER := env.Get("DB_USER")
	DB_NAME := env.Get("DB_NAME")
	DB_PASSWORD := env.Get("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_NAME, DB_PASSWORD)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	logger.ErrorWithMessage("Error connecting to database", err)

	sqlDB, err := DB.DB()
	logger.ErrorWithMessage("Error getting database connection", err)

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	logger.Info("Database connection initialized")
}
