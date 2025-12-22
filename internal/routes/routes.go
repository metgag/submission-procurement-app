package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func InitRoutes(app *fiber.App) {
	app.Use(logger.New())

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
}
