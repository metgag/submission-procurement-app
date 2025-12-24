package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/metgag/procurement-api-example/internal/dto"
	"github.com/metgag/procurement-api-example/internal/models"
	"github.com/metgag/procurement-api-example/internal/services"
	"github.com/metgag/procurement-api-example/internal/utils"
)

type SupplierItemHandler struct {
	Service *services.SupplierItemService
}

func NewSupplierItemHandler(service *services.SupplierItemService) *SupplierItemHandler {
	return &SupplierItemHandler{Service: service}
}

func (h *SupplierItemHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateSupplierItemRequest
	if err := utils.ParseBodyAndValidate(c, &req); err != nil {
		return err
	}

	// Validasi existence Supplier & Item
	var supplier models.Supplier
	if err := utils.FindByIDValue(c, h.Service.DB, "Supplier", req.SupplierID, &supplier); err != nil {
		return err
	}
	var item models.Item
	if err := utils.FindByIDValue(c, h.Service.DB, "Item", req.ItemID, &item); err != nil {
		return err
	}

	supplierItem, err := h.Service.Create(req)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Unable to assign item to supplier", err, true)
	}

	return c.Status(fiber.StatusCreated).JSON(dto.OKResponse{
		Message: "Item assigned to supplier successfully",
		Data: dto.SupplierItemResponse{
			ID:           supplierItem.ID,
			SupplierID:   supplierItem.SupplierID,
			ItemID:       supplierItem.ItemID,
			Price:        supplierItem.Price,
			Stock:        supplierItem.Stock,
			ItemName:     item.Name,
			SupplierName: supplier.Name,
		},
	})
}

func (h *SupplierItemHandler) ReadItems(c *fiber.Ctx) error {
	supplierID := c.QueryInt("supplier_id")
	items, err := h.Service.GetAll(supplierID)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Unable to retrieve supplier items", err, true)
	}

	res := make([]dto.SupplierItemResponse, 0, len(items))
	for _, si := range items {
		res = append(res, dto.SupplierItemResponse{
			ID:           si.ID,
			SupplierID:   si.SupplierID,
			ItemID:       si.ItemID,
			Price:        si.Price,
			Stock:        si.Stock,
			ItemName:     si.Item.Name,
			SupplierName: si.Supplier.Name,
		})
	}

	return c.JSON(dto.OKResponse{
		Message: "Supplier items retrieved successfully",
		Data:    res,
	})
}

func (h *SupplierItemHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var supplierItem models.SupplierItem
	if err := utils.FindByIDValue(c, h.Service.DB, "SupplierItem", id, &supplierItem); err != nil {
		return err
	}

	var req dto.UpdateSupplierItemRequest
	if err := utils.ParseBodyAndValidate(c, &req); err != nil {
		return err
	}

	if err := h.Service.Update(&supplierItem, req); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Failed to update supplier item", err, true)
	}

	return c.JSON(dto.OKResponse{
		Message: "Supplier item updated successfully",
		Data: dto.SupplierItemResponse{
			ID:         supplierItem.ID,
			SupplierID: supplierItem.SupplierID,
			ItemID:     supplierItem.ItemID,
			Price:      supplierItem.Price,
			Stock:      supplierItem.Stock,
		},
	})
}

func (h *SupplierItemHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	var supplierItem models.SupplierItem
	if err := utils.FindByIDValue(c, h.Service.DB, "SupplierItem", id, &supplierItem); err != nil {
		return err
	}

	if err := h.Service.Delete(&supplierItem); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Failed to delete supplier item", err, true)
	}

	return c.JSON(dto.OKResponse{
		Message: "Supplier item deleted successfully",
	})
}
