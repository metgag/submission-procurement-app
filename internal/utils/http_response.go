package utils

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/metgag/procurement-api-example/internal/dto"
)

func Error(
	c *fiber.Ctx,
	status int,
	publicMessage string,
	err error,
	logAsError bool,
) error {
	if err != nil && logAsError {
		log.Printf("[ERROR] %s: %v\n", publicMessage, err)
	}

	_ = c.Status(status).JSON(dto.ErrorResponse{
		Message: publicMessage,
	})

	return fiber.NewError(status, publicMessage)
}
