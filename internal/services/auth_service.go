package services

import (
	"errors"

	"gorm.io/gorm"

	"github.com/metgag/procurement-api-example/internal/dto"
	"github.com/metgag/procurement-api-example/internal/models"
	"github.com/metgag/procurement-api-example/internal/utils"
)

type AuthService struct {
	DB *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{DB: db}
}

func (s *AuthService) Register(req dto.RegisterRequest) (*dto.RegisterResponse, error) {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Username: req.Username,
		Password: hashedPassword,
		Role:     models.UserRole(req.Role),
	}

	if err := s.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &dto.RegisterResponse{
		ID:       user.ID,
		Username: user.Username,
		Role:     string(user.Role),
	}, nil
}

func (s *AuthService) Login(req dto.LoginRequest) (string, error) {
	var user models.User

	if err := s.DB.
		Where("username = ?", req.Username).
		First(&user).
		Error; err != nil {
		return "", errors.New("invalid username or password")
	}

	if err := utils.ComparePassword(user.Password, req.Password); err != nil {
		return "", errors.New("invalid username or password")
	}

	token, err := utils.GenerateJWT(user.ID, user.Username, string(user.Role))
	if err != nil {
		return "", err
	}

	return token, nil
}
