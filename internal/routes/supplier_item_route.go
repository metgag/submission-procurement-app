package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/metgag/procurement-api-example/internal/handlers"
	"github.com/metgag/procurement-api-example/internal/middleware"
)

func RegisterSupplierItemRoute(api fiber.Router, supplierItemHandler *handlers.SupplierItemHandler) {
	authProtected := api.Group("/", middleware.JWTProtected())
	authProtected.Get("/supplier-items", supplierItemHandler.ReadItems)

	admin := api.Group("/admin", middleware.JWTProtected(), middleware.RequireRole("admin"))
	adminSupplierItems := admin.Group("/supplier-items")
	{
		adminSupplierItems.Post("/", supplierItemHandler.Create)
		adminSupplierItems.Patch("/:id", supplierItemHandler.Update)
		adminSupplierItems.Delete("/:id", supplierItemHandler.Delete)
	}
}
