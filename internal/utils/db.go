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

func FindByIDValue(
	c *fiber.Ctx,
	db *gorm.DB,
	modelName string,
	id any,
	dest any,
) error {
	if err := db.First(dest, id).Error; err != nil {
		return Error(c, fiber.StatusNotFound, modelName+" not found", err, true)
	}
	return nil
}

func GetPaginationParams(c *fiber.Ctx) (page int, pageSize int) {
	page = c.QueryInt("page", 1)
	if page < 1 {
		page = 1
	}

	pageSize = c.QueryInt("page_size", 10)
	if pageSize < 1 {
		pageSize = 10
	}

	return
}
