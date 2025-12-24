package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/metgag/procurement-api-example/internal/handlers"
	"github.com/metgag/procurement-api-example/internal/middleware"
)

func RegisterItemRoute(api fiber.Router, itemHandler *handlers.ItemHandler) {
	authProtected := api.Group("/", middleware.JWTProtected())
	authProtected.Get("/items", itemHandler.ReadAll)

	admin := api.Group("/admin", middleware.JWTProtected(), middleware.RequireRole("admin"))
	adminItems := admin.Group("/items")
	{
		adminItems.Post("/", itemHandler.Create)
		adminItems.Patch("/:id", itemHandler.Update)
		adminItems.Delete("/:id", itemHandler.Delete)
	}
}
