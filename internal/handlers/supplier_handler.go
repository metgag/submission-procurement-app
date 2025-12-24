package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/metgag/procurement-api-example/internal/dto"
	"github.com/metgag/procurement-api-example/internal/models"
	"github.com/metgag/procurement-api-example/internal/services"
	"github.com/metgag/procurement-api-example/internal/utils"
)

type SupplierHandler struct {
	Service *services.SupplierService
}

func NewSupplierHandler(service *services.SupplierService) *SupplierHandler {
	return &SupplierHandler{Service: service}
}

func (h *SupplierHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateSupplierRequest
	if err := utils.ParseBodyAndValidate(c, &req); err != nil {
		return err
	}

	supplier, err := h.Service.Create(req)
	if err != nil {
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
	page, pageSize := utils.GetPaginationParams(c)

	suppliers, total, err := h.Service.GetAll(page, pageSize)
	if err != nil {
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
		Meta: &dto.MetaResponse{
			Page:       page,
			PageSize:   pageSize,
			TotalItems: total,
			TotalPages: (total + int64(pageSize) - 1) / int64(pageSize),
		},
	})
}

func (h *SupplierHandler) Update(c *fiber.Ctx) error {
	var supplier models.Supplier
	if err := utils.FindByID(c, h.Service.DB, "Supplier", &supplier); err != nil {
		return err
	}

	var req dto.UpdateSupplierRequest
	if err := utils.ParseBodyAndValidate(c, &req); err != nil {
		return err
	}

	if err := h.Service.Update(&supplier, req); err != nil {
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
	if err := utils.FindByID(c, h.Service.DB, "Supplier", &supplier); err != nil {
		return err
	}

	if err := h.Service.Delete(&supplier); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Failed to delete supplier", err, true)
	}

	log.Printf("Supplier deleted: id=%d", supplier.ID)
	return c.JSON(dto.OKResponse{
		Message: "Supplier deleted successfully",
	})
}
