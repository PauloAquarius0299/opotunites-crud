package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

// Init initializes the PostgreSQL database connection
func Init() error {
	logger = GetLogger("config") // Initialize logger for the config package

	var err error
	db, err = InitializePostgreSQL()
	if err != nil {
		logger.Errorf("Error initializing database: %v", err)
		return fmt.Errorf("error initializing PostgreSQL: %w", err)
	}

	return nil
}

// GetDB returns the PostgreSQL database instance
func GetDB() *gorm.DB {
	return db
}

// GetLogger returns a logger instance
func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}