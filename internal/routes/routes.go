package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"

	"github.com/metgag/procurement-api-example/internal/handlers"
	"github.com/metgag/procurement-api-example/internal/middleware"
	"github.com/metgag/procurement-api-example/internal/services"
)

func InitRoutes(app *fiber.App, db *gorm.DB) {
	app.Use(logger.New())

	// ========================
	// Handlers & Services
	// ========================
	authHandler := handlers.NewAuthHandler(db)
	supplierHandler := handlers.NewSupplierHandler(db)
	itemHandler := handlers.NewItemHandler(db)
	supplierItemHandler := handlers.NewSupplierItemHandler(db)

	purchaseService := services.NewPurchasingService(db)
	purchaseHandler := handlers.NewPurchasingHandler(purchaseService)

	// ========================
	// API Group
	// ========================
	api := app.Group("/api")
	{
		// ========================
		// Public Routes
		// ========================
		api.Get("/ping", func(c *fiber.Ctx) error {
			return c.SendString("pong")
		})

		// ========================
		// Auth Routes
		// ========================
		auth := api.Group("/auth")
		auth.Post("/register", authHandler.Register)
		auth.Post("/login", authHandler.Login)

		// ========================
		// Admin Routes (protected)
		// ========================
		admin := api.Group(
			"/admin",
			middleware.JWTProtected(),
			middleware.RequireRole("admin"),
		)
		{
			adminSuppliers := admin.Group("/suppliers")
			{
				adminSuppliers.Post("/", supplierHandler.Create)
				adminSuppliers.Get("/", supplierHandler.ReadAll)
				adminSuppliers.Patch("/:id", supplierHandler.Update)
				adminSuppliers.Delete("/:id", supplierHandler.Delete)
			}

			adminItems := admin.Group("/items")
			{
				adminItems.Post("/", itemHandler.Create)
				adminItems.Get("/", itemHandler.ReadAll)
			}

			adminSupplierItems := admin.Group("/supplier-items")
			{
				adminSupplierItems.Post("/", supplierItemHandler.Create)
				adminSupplierItems.Get("/", supplierItemHandler.ReadItemsBySupplierID)
			}
		}

		// ========================
		// User Routes (protected)
		// ========================
		user := api.Group(
			"/",
			middleware.JWTProtected(),
			middleware.RequireRole("user"),
		)
		{
			userPurchases := user.Group("/purchases")
			{
				userPurchases.Post("/", purchaseHandler.Create)
			}
		}
	}
}
