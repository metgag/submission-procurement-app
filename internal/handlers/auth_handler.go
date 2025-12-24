package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/metgag/procurement-api-example/internal/dto"
	"github.com/metgag/procurement-api-example/internal/models"
	"github.com/metgag/procurement-api-example/internal/utils"
)

type AuthHandler struct {
	DB *gorm.DB
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{DB: db}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req dto.RegisterRequest

	if err := utils.ParseAndValidate(c, &req); err != nil {
		return err
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Failed to hash password", err, true)
	}

	user := models.User{
		Username: req.Username,
		Password: hashedPassword,
		Role:     models.UserRole(req.Role),
	}

	if err := h.DB.Create(&user).Error; err != nil {
		return utils.Error(c, fiber.StatusConflict, "Username already taken", err, true)
	}

	return c.Status(fiber.StatusCreated).JSON(dto.OKResponse{
		Message: "User registered successfully",
		Data: dto.RegisterResponse{
			ID:       user.ID,
			Username: user.Username,
			Role:     string(user.Role),
		},
	})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req dto.LoginRequest

	if err := utils.ParseAndValidate(c, &req); err != nil {
		return err
	}

	var user models.User
	if err := h.DB.
		Where("username = ?", req.Username).
		First(&user).
		Error; err != nil {
		return utils.Error(c, fiber.StatusUnauthorized, "Invalid username or password", err, false)
	}

	if err := utils.ComparePassword(user.Password, req.Password); err != nil {
		return utils.Error(c, fiber.StatusUnauthorized, "Invalid username or password", err, false)
	}

	token, err := utils.GenerateJWT(user.ID, user.Username, string(user.Role))
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Failed to generate token", err, true)
	}

	return c.JSON(dto.OKResponse{
		Message: "Login successfully",
		Data: dto.LoginResponse{
			Token: token,
		},
	})
}
