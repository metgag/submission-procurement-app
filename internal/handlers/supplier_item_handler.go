package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/metgag/procurement-api-example/internal/dto"
	"github.com/metgag/procurement-api-example/internal/models"
	"github.com/metgag/procurement-api-example/internal/utils"
)

type SupplierItemHandler struct {
	DB *gorm.DB
}

func NewSupplierItemHandler(db *gorm.DB) *SupplierItemHandler {
	return &SupplierItemHandler{DB: db}
}

func (h *SupplierItemHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateSupplierItemRequest
	if err := utils.ParseAndValidate(c, &req); err != nil {
		return err
	}

	var supplier models.Supplier
	if err := utils.FindByIDValue(c, h.DB, "Supplier", req.SupplierID, &supplier); err != nil {
		return err
	}

	var item models.Item
	if err := utils.FindByIDValue(c, h.DB, "Item", req.ItemID, &item); err != nil {
		return err
	}

	supplierItem := models.SupplierItem{
		SupplierID: req.SupplierID,
		ItemID:     req.ItemID,
		Price:      req.Price,
		Stock:      req.Stock,
	}

	if err := h.DB.Create(&supplierItem).Error; err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Unable to assign item to supplier", err, true)
	}

	return c.Status(fiber.StatusCreated).JSON(dto.OKResponse{
		Message: "Item assigned to supplier successfully",
		Data: dto.SupplierItemResponse{
			ID:         supplierItem.ID,
			SupplierID: supplierItem.SupplierID,
			ItemID:     supplierItem.ItemID,
			Price:      supplierItem.Price,
			Stock:      supplierItem.Stock,
		},
	})
}

func (h *SupplierItemHandler) ReadItems(c *fiber.Ctx) error {
	var items []models.SupplierItem

	query := h.DB.
		Preload("Item").
		Preload("Supplier")

	if supplierID := c.QueryInt("supplier_id"); supplierID > 0 {
		query = query.Where("supplier_id = ?", supplierID)
	}

	if err := query.Find(&items).Error; err != nil {
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
		Message: "Supplier items retrieved succesfully",
		Data:    res,
	})
}

func (h *SupplierItemHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	var supplierItem models.SupplierItem
	if err := utils.FindByIDValue(c, h.DB, "SupplierItem", id, &supplierItem); err != nil {
		return err
	}

	var req dto.UpdateSupplierItemRequest
	if err := utils.ParseAndValidate(c, &req); err != nil {
		return err
	}

	if req.Price != nil {
		supplierItem.Price = *req.Price
	}
	if req.Stock != nil {
		supplierItem.Stock = *req.Stock
	}

	if err := h.DB.Save(&supplierItem).Error; err != nil {
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
	if err := utils.FindByIDValue(c, h.DB, "SupplierItem", id, &supplierItem); err != nil {
		return err
	}

	if err := h.DB.Delete(&supplierItem).Error; err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Failed to delete supplier item", err, true)
	}

	return c.JSON(dto.OKResponse{
		Message: "Supplier item deleted successfully",
	})
}
