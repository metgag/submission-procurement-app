package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/metgag/procurement-api-example/internal/handlers"
	"github.com/metgag/procurement-api-example/internal/middleware"
)

func RegisterSupplierRoute(api fiber.Router, supplierHandler *handlers.SupplierHandler) {
	authProtected := api.Group("/", middleware.JWTProtected())
	authProtected.Get("/suppliers", supplierHandler.ReadAll)

	admin := api.Group("/admin", middleware.JWTProtected(), middleware.RequireRole("admin"))
	adminSuppliers := admin.Group("/suppliers")
	{
		adminSuppliers.Post("/", supplierHandler.Create)
		adminSuppliers.Patch("/:id", supplierHandler.Update)
		adminSuppliers.Delete("/:id", supplierHandler.Delete)
	}
}
