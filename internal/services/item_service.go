package services

import (
	"errors"

	"gorm.io/gorm"

	"github.com/metgag/procurement-api-example/internal/dto"
	"github.com/metgag/procurement-api-example/internal/models"
)

type ItemService struct {
	DB *gorm.DB
}

func NewItemService(db *gorm.DB) *ItemService {
	return &ItemService{DB: db}
}

func (s *ItemService) Create(req dto.CreateItemRequest) (*models.Item, error) {
	item := models.Item{
		Name: req.Name,
	}

	if err := s.DB.Create(&item).Error; err != nil {
		return nil, err
	}

	return &item, nil
}

func (s *ItemService) GetAll() ([]models.Item, error) {
	var items []models.Item

	if err := s.DB.Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (s *ItemService) Update(item *models.Item, req dto.CreateItemRequest) error {
	item.Name = req.Name

	if err := s.DB.Save(item).Error; err != nil {
		return err
	}

	return nil
}

func (s *ItemService) Delete(item *models.Item) error {
	if err := s.DB.Delete(item).Error; err != nil {
		return err
	}
	return nil
}

// Optional: business validation example
func (s *ItemService) ValidateCreate(req dto.CreateItemRequest) error {
	if req.Name == "" {
		return errors.New("item name cannot be empty")
	}
	return nil
}
