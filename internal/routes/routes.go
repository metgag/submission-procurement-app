package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"

	"github.com/metgag/procurement-api-example/internal/handlers"
	"github.com/metgag/procurement-api-example/internal/services"
)

func InitRoutes(app *fiber.App, db *gorm.DB) {
	app.Use(logger.New())

	// =========================
	// Static Frontend
	// =========================
	app.Get("/", func(c *fiber.Ctx) error { return c.Redirect("/login.html") })
	app.Get("/login", func(c *fiber.Ctx) error { return c.SendFile("./client/login.html") })
	app.Get("/register", func(c *fiber.Ctx) error { return c.SendFile("./client/register.html") })
	app.Get("/dashboard", func(c *fiber.Ctx) error { return c.SendFile("./client/dashboard.html") })
	app.Get("/purchase", func(c *fiber.Ctx) error { return c.SendFile("./client/purchase.html") })
	app.Get("/invoice", func(c *fiber.Ctx) error { return c.SendFile("./client/invoice.html") })

	app.Static("/", "./client")

	// =========================
	// Initialize Services & Handlers
	// =========================
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

	// =========================
	// API Group
	// =========================
	api := app.Group("/api/v1")
	api.Get("/ping", func(c *fiber.Ctx) error { return c.SendString("pong") })

	// =========================
	// Register routes modularly
	// =========================
	RegisterAuthRoute(api, authHandler)
	RegisterSupplierRoute(api, supplierHandler)
	RegisterItemRoute(api, itemHandler)
	RegisterSupplierItemRoute(api, supplierItemHandler)
	RegisterPurchaseRoute(api, purchaseHandler)
}
