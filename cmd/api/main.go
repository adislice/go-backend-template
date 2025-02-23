package main

import (
	"github.com/adislice/go-project-structure/config"
	"github.com/adislice/go-project-structure/internal/database"
	"github.com/adislice/go-project-structure/internal/routes"
	"github.com/adislice/go-project-structure/pkg/logger"
	"github.com/adislice/go-project-structure/pkg/validation"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load konfigurasi dari environment variable
	config.LoadConfig()

	// Koneksi ke database
	db := database.ConnectDB()
	db = db.Debug()

	// Insiasi validator
	validator := validation.NewValidator()

	// Inisiasi logger
	log := logger.InitLogger()
	defer log.Sync()

	// Inisasi fiber
	app := fiber.New()

	// Setup routes
	routes.ApiRoutes(app, db, validator)

	// Jalankan server
	port := config.AppConfig.AppPort
	if port == "" {
		port = "7000"
	}

	app.Listen(":" + port)

}
