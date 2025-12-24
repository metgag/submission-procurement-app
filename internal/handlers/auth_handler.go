package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/metgag/procurement-api-example/internal/dto"
	"github.com/metgag/procurement-api-example/internal/services"
	"github.com/metgag/procurement-api-example/internal/utils"
)

type AuthHandler struct {
	AuthService *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: service}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req dto.RegisterRequest

	if err := utils.ParseBodyAndValidate(c, &req); err != nil {
		return err
	}

	res, err := h.AuthService.Register(req)
	if err != nil {
		return utils.Error(
			c,
			fiber.StatusConflict,
			"Username already taken",
			err,
			true,
		)
	}

	return c.Status(fiber.StatusCreated).JSON(dto.OKResponse{
		Message: "User registered successfully",
		Data:    res,
	})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req dto.LoginRequest

	if err := utils.ParseBodyAndValidate(c, &req); err != nil {
		return err
	}

	token, err := h.AuthService.Login(req)
	if err != nil {
		return utils.Error(
			c,
			fiber.StatusUnauthorized,
			"Invalid username or password",
			err,
			false,
		)
	}

	return c.JSON(dto.OKResponse{
		Message: "Login successfully",
		Data: dto.LoginResponse{
			Token: token,
		},
	})
}
