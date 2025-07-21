package config

import (
	"crud-oportunides/schemas"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func InitializePostgreSQL() (*gorm.DB, error) {
	logger := GetLogger("postgresql")
	_ = godotenv.Load()

	// Connection string for PostgreSQL
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	// Connect to PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error connecting to PostgreSQL: %v", err)
		return nil, err
	}

	// Verify connection
	sqlDB, err := db.DB()
	if err != nil {
		logger.Errorf("Error getting DB instance: %v", err)
		return nil, err
	}

	err = sqlDB.Ping()
	if err != nil {
		logger.Errorf("Error pinging database: %v", err)
		return nil, err
	}

	// Run migrations
	err = db.AutoMigrate(&schemas.OpeningResponse{})
	if err != nil {
		logger.Errorf("Error migrating database: %v", err)
		return nil, err
	}

	logger.Info("Successfully connected to PostgreSQL")
	return db, nil
}