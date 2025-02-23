package database

import (
	"fmt"
	"log"

	"github.com/adislice/go-project-structure/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dbHost := config.AppConfig.DBHost
	dbPort := config.AppConfig.DBPort
	dbUser := config.AppConfig.DBUsername
	dbPass := config.AppConfig.DBPassword
	dbName := config.AppConfig.DBName

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return db
}
