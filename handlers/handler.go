package handlers

import (
	"crud-oportunides/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handlers")
	var err error
	db, err = config.InitializePostgreSQL() // Substitui GetSQLite por InitializePostgreSQL
	if err != nil {
		logger.Errorf("Failed to initialize PostgreSQL: %v", err)
		panic(err) // Ou trate o erro de forma adequada
	}
}