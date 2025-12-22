package handlers

import (
	"log"

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

func (h *SupplierHandler) ReadAll(c *fiber.Ctx) error {
	var suppliers []models.Supplier
	if err := h.DB.Find(&suppliers).Error; err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Failed to read suppliers", err, true)
	}

	res := make([]dto.SupplierResponse, 0, len(suppliers))
	for _, s := range suppliers {
		res = append(res, dto.SupplierResponse{
			ID:      s.ID,
			Name:    s.Name,
			Email:   s.Email,
			Address: s.Address,
		})
	}

	return c.JSON(dto.OKResponse{
		Message: "Suppliers retrieved successfully",
		Data:    res,
	})
}

func (h *SupplierHandler) Update(c *fiber.Ctx) error {
	var supplier models.Supplier
	if err := utils.FindByID(c, h.DB, "Supplier", &supplier); err != nil {
		return err
	}

	var req dto.UpdateSupplierRequest
	if err := utils.ParseAndValidate(c, &req); err != nil {
		return err
	}

	if req.Name != nil {
		supplier.Name = *req.Name
	}
	if req.Email != nil {
		supplier.Email = *req.Email
	}
	if req.Address != nil {
		supplier.Address = *req.Address
	}

	if err := h.DB.Save(&supplier).Error; err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Failed to update supplier", err, true)
	}

	log.Printf("Supplier updated: id=%d", supplier.ID)
	return c.JSON(dto.OKResponse{
		Message: "Supplier updated successfully",
		Data: dto.SupplierResponse{
			ID:      supplier.ID,
			Name:    supplier.Name,
			Email:   supplier.Email,
			Address: supplier.Address,
		},
	})
}

func (h *SupplierHandler) Delete(c *fiber.Ctx) error {
	var supplier models.Supplier
	if err := utils.FindByID(c, h.DB, "Supplier", &supplier); err != nil {
		return err
	}

	if err := h.DB.Delete(&supplier).Error; err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Failed to delete supplier", err, true)
	}

	log.Printf("Supplier deleted: id=%d", supplier.ID)
	return c.JSON(dto.OKResponse{
		Message: "Supplier deleted successfully",
	})
}
