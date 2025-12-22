package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"github.com/metgag/procurement-api-example/internal/dto"
	"github.com/metgag/procurement-api-example/internal/utils"
)

func RequireRole(role string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)

		claims, ok := user.Claims.(*utils.JWTClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{
				Message: "Token claims unrecognized",
			})
		}

		if claims.Role != role {
			return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{
				Message: "Access role denied",
			})
		}

		return c.Next()
	}
}
