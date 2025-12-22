package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/metgag/procurement-api-example/internal/dto"
	"github.com/metgag/procurement-api-example/internal/models"
	"github.com/metgag/procurement-api-example/internal/utils"
)

type SupplierHandler struct {
	DB *gorm.DB
}

func NewSupplierHandler(db *gorm.DB) *SupplierHandler {
	return &SupplierHandler{DB: db}
}

func (h *SupplierHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateSupplierRequest
	if err := utils.ParseAndValidate(c, &req); err != nil {
		return err
	}

	supplier := models.Supplier{
		Name:    req.Name,
		Email:   req.Email,
		Address: req.Address,
	}

	if err := h.DB.
		Create(&supplier).
		Error; err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Failed to create supplier", err, true)
	}

	return c.Status(fiber.StatusCreated).JSON(dto.OKResponse{
		Message: "Supplier created successfully",
		Data: dto.SupplierResponse{
			ID:      supplier.ID,
			Name:    supplier.Name,
			Email:   supplier.Email,
			Address: supplier.Address,
		},
	})
}
