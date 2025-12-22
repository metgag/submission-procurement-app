package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/metgag/procurement-api-example/internal/config"
	"github.com/metgag/procurement-api-example/internal/models"
	"github.com/metgag/procurement-api-example/internal/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.InitDB()
	if err := db.AutoMigrate(
		&models.User{},
		&models.Supplier{},
	); err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}

	app := fiber.New()
	routes.InitRoutes(app, db)

	app.Listen(":3080")
}
