package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/metgag/procurement-api-example/internal/handlers"
	"github.com/metgag/procurement-api-example/internal/middleware"
)

func RegisterPurchaseRoute(api fiber.Router, purchaseHandler *handlers.PurchasingHandler) {
	user := api.Group("/", middleware.JWTProtected(), middleware.RequireRole("user"))
	userPurchases := user.Group("/purchases")
	userPurchases.Post("/", purchaseHandler.Create)
}
