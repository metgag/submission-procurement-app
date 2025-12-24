package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/metgag/procurement-api-example/internal/dto"
	"github.com/metgag/procurement-api-example/internal/models"
	"github.com/metgag/procurement-api-example/internal/services"
	"github.com/metgag/procurement-api-example/internal/utils"
)

type ItemHandler struct {
	ItemService *services.ItemService
}

func NewItemHandler(service *services.ItemService) *ItemHandler {
	return &ItemHandler{ItemService: service}
}

func (h *ItemHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateItemRequest
	if err := utils.ParseBodyAndValidate(c, &req); err != nil {
		return err
	}

	item, err := h.ItemService.Create(req)
	if err != nil {
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
	items, err := h.ItemService.GetAll()
	if err != nil {
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
	if err := utils.ParseBodyAndValidate(c, &req); err != nil {
		return err
	}

	var item models.Item
	if err := utils.FindByID(c, h.ItemService.DB, "Item", &item); err != nil {
		return err
	}

	if err := h.ItemService.Update(&item, req); err != nil {
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
	if err := utils.FindByID(c, h.ItemService.DB, "Item", &item); err != nil {
		return err
	}

	if err := h.ItemService.Delete(&item); err != nil {
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
