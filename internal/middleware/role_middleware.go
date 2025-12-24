package middleware

import (
	"slices"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"github.com/metgag/procurement-api-example/internal/dto"
	"github.com/metgag/procurement-api-example/internal/utils"
)

func RequireRole(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user, ok := c.Locals("user").(*jwt.Token)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{
				Message: "Invalid or missing token",
			})
		}

		claims, ok := user.Claims.(*utils.JWTClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{
				Message: "Token claims unrecognized",
			})
		}

		if slices.Contains(allowedRoles, claims.Role) {
			return c.Next()
		}

		return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{
			Message: "Access role denied",
		})
	}
}
