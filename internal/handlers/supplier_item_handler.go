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

func (h *SupplierItemHandler) ReadItemsBySupplierID(c *fiber.Ctx) error {
	supplierID := c.QueryInt("supplier_id")
	if supplierID == 0 {
		return utils.Error(c, fiber.StatusBadRequest, "supplier_id is required", nil, false)
	}

	var items []models.SupplierItem
	if err := h.DB.
		Preload("Item").
		Preload("Supplier").
		Where("supplier_id = ?", supplierID).
		Find(&items).
		Error; err != nil {
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
		Message: "Items by supplier retrieved successfully",
		Data:    res,
	})
}
