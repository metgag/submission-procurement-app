package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/metgag/procurement-api-example/internal/dto"
	"github.com/metgag/procurement-api-example/internal/services"
	"github.com/metgag/procurement-api-example/internal/utils"
)

type PurchasingHandler struct {
	Service *services.PurchasingService
}

func NewPurchasingHandler(s *services.PurchasingService) *PurchasingHandler {
	return &PurchasingHandler{Service: s}
}

func (h *PurchasingHandler) Create(c *fiber.Ctx) error {
	var req dto.CreatePurchasingRequest
	if err := utils.ParseBodyAndValidate(c, &req); err != nil {
		return err
	}

	claims := utils.GetJWTClaims(c)
	purchase, items, err := h.Service.Create(claims.UserID, req)
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Failed to create purchase", err, true)
	}

	return c.Status(fiber.StatusCreated).JSON(dto.OKResponse{
		Message: "Purchase created successfully",
		Data: dto.PurchasingResponse{
			ID:         purchase.ID,
			GrandTotal: purchase.GrandTotal,
			Items:      items,
		},
	})
}
