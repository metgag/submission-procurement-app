package utils

import "github.com/gofiber/fiber/v2"

func ParseAndValidate(c *fiber.Ctx, req any) error {
	if err := c.BodyParser(req); err != nil {
		return Error(c, fiber.StatusBadRequest, "Invalid request body", err, true)
	}

	if err := Validate.Struct(req); err != nil {
		return Error(c, fiber.StatusBadRequest, "Validation failed", err, true)
	}

	return nil
}
