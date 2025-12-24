package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/metgag/procurement-api-example/internal/handlers"
)

func RegisterAuthRoute(api fiber.Router, authHandler *handlers.AuthHandler) {
	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)
}
