package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"

	"github.com/metgag/procurement-api-example/internal/handlers"
)

func InitRoutes(app *fiber.App, db *gorm.DB) {
	app.Use(logger.New())

	authHandler := handlers.NewAuthHandler(db)

	api := app.Group("/api")
	{
		api.Get("/ping", func(c *fiber.Ctx) error {
			return c.SendString("pong")
		})

		auth := api.Group("/auth")
		{
			auth.Post("/register", authHandler.Register)
			auth.Post("/login", authHandler.Login)
		}
	}
}
