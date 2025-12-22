package utils

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func FindByID(c *fiber.Ctx, db *gorm.DB, modelName string, dest any) error {
	id := c.Params("id")
	if err := db.First(dest, id).Error; err != nil {
		return Error(c, fiber.StatusNotFound, modelName+" not found", err, true)
	}

	return nil
}
