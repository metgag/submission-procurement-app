package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"

	"github.com/metgag/procurement-api-example/internal/handlers"
	"github.com/metgag/procurement-api-example/internal/middleware"
)

func InitRoutes(app *fiber.App, db *gorm.DB) {
	app.Use(logger.New())

	authHandler := handlers.NewAuthHandler(db)
	supplierHandler := handlers.NewSupplierHandler(db)

	api := app.Group("/api")

	// public routes
	api.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	// auth routes
	auth := api.Group("/auth")
	{
		auth.Post("/register", authHandler.Register)
		auth.Post("/login", authHandler.Login)
	}

	// protected admin routes
	admin := api.Group(
		"/admin",
		middleware.JWTProtected(),
		middleware.RequireRole("admin"),
	)
	{
		admin.Post("/suppliers", supplierHandler.Create)
	}
}
