package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/metgag/procurement-api-example/internal/dto"
	"github.com/metgag/procurement-api-example/internal/models"
	"github.com/metgag/procurement-api-example/internal/utils"
)

type ItemHandler struct {
	DB *gorm.DB
}

func NewItemHandler(db *gorm.DB) *ItemHandler {
	return &ItemHandler{DB: db}
}

func (h *ItemHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateItemRequest
	if err := utils.ParseAndValidate(c, &req); err != nil {
		return err
	}

	item := models.Item{Name: req.Name}
	if err := h.DB.Create(&item).Error; err != nil {
		return utils.Error(c, fiber.StatusConflict, "Item already exist", err, true)
	}

	return c.Status(fiber.StatusCreated).JSON(dto.OKResponse{
		Message: "Item created successfully",
		Data: dto.ItemResponse{
			ID:   item.ID,
			Name: item.Name,
		},
	})
}

func (h *ItemHandler) ReadAll(c *fiber.Ctx) error {
	var items []models.Item
	if err := h.DB.Find(&items).Error; err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Unable to retrieve items", err, true)
	}

	res := make([]dto.ItemResponse, 0, len(items))
	for _, item := range items {
		res = append(res, dto.ItemResponse{
			ID:   item.ID,
			Name: item.Name,
		})
	}

	return c.JSON(dto.OKResponse{
		Message: "Items retrieved successfully",
		Data:    res,
	})
}

func (h *ItemHandler) Update(c *fiber.Ctx) error {
	var req dto.CreateItemRequest
	if err := utils.ParseAndValidate(c, &req); err != nil {
		return err
	}

	var item models.Item
	if err := utils.FindByID(c, h.DB, "Item", &item); err != nil {
		return err
	}

	item.Name = req.Name

	if err := h.DB.Save(&item).Error; err != nil {
		return utils.Error(
			c,
			fiber.StatusInternalServerError,
			"Unable to update item",
			err,
			true,
		)
	}

	return c.JSON(dto.OKResponse{
		Message: "Item updated successfully",
		Data: dto.ItemResponse{
			ID:   item.ID,
			Name: item.Name,
		},
	})
}

func (h *ItemHandler) Delete(c *fiber.Ctx) error {
	var item models.Item
	if err := utils.FindByID(c, h.DB, "Item", &item); err != nil {
		return err
	}

	if err := h.DB.Delete(&item).Error; err != nil {
		return utils.Error(
			c,
			fiber.StatusInternalServerError,
			"Unable to delete item",
			err,
			true,
		)
	}

	return c.JSON(dto.OKResponse{
		Message: "Item deleted successfully",
		Data: dto.ItemResponse{
			ID:   item.ID,
			Name: item.Name,
		},
	})
}
