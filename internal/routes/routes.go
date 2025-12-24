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

	// =========================
	// Static Frontend
	// =========================

	// Redirect root ke login
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/login.html")
	})

	// Optional clean URLs
	app.Get("/login", func(c *fiber.Ctx) error {
		return c.SendFile("./client/login.html")
	})
	app.Get("/register", func(c *fiber.Ctx) error {
		return c.SendFile("./client/register.html")
	})
	app.Get("/dashboard", func(c *fiber.Ctx) error {
		return c.SendFile("./client/dashboard.html")
	})
	app.Get("/purchase", func(c *fiber.Ctx) error {
		return c.SendFile("./client/purchase.html")
	})
	app.Get("/invoice", func(c *fiber.Ctx) error {
		return c.SendFile("./client/invoice.html")
	})

	app.Static("/", "./client")

	// ========================
	// Handlers & Services
	// ========================
	authService := services.NewAuthService(db)
	authHandler := handlers.NewAuthHandler(authService)

	supplierService := services.NewSupplierService(db)
	supplierHandler := handlers.NewSupplierHandler(supplierService)

	itemService := services.NewItemService(db)
	itemHandler := handlers.NewItemHandler(itemService)

	supplierItemService := services.NewSupplierItemService(db)
	supplierItemHandler := handlers.NewSupplierItemHandler(supplierItemService)

	purchaseService := services.NewPurchasingService(db)
	purchaseHandler := handlers.NewPurchasingHandler(purchaseService)

	// ========================
	// API Group
	// ========================
	api := app.Group("/api/v1")
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
		// Admin & User Routes (protected)
		// ========================
		authProtected := api.Group(
			"/",
			middleware.JWTProtected(),
		)
		{
			authProtected.Get("/items", itemHandler.ReadAll)
			authProtected.Get("/suppliers", supplierHandler.ReadAll)
			authProtected.Get("/supplier-items", supplierItemHandler.ReadItems)
		}

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
				adminSuppliers.Patch("/:id", supplierHandler.Update)
				adminSuppliers.Delete("/:id", supplierHandler.Delete)
			}

			adminItems := admin.Group("/items")
			{
				adminItems.Post("/", itemHandler.Create)
				adminItems.Patch("/:id", itemHandler.Update)
				adminItems.Delete("/:id", itemHandler.Delete)
			}

			adminSupplierItems := admin.Group("/supplier-items")
			{
				adminSupplierItems.Post("/", supplierItemHandler.Create)
				adminSupplierItems.Patch("/:id", supplierItemHandler.Update)
				adminSupplierItems.Delete("/:id", supplierItemHandler.Delete)
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
