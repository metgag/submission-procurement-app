package services

import (
	"gorm.io/gorm"

	"github.com/metgag/procurement-api-example/internal/dto"
	"github.com/metgag/procurement-api-example/internal/models"
)

type SupplierService struct {
	DB *gorm.DB
}

func NewSupplierService(db *gorm.DB) *SupplierService {
	return &SupplierService{DB: db}
}

func (s *SupplierService) Create(req dto.CreateSupplierRequest) (*models.Supplier, error) {
	supplier := models.Supplier{
		Name:    req.Name,
		Email:   req.Email,
		Address: req.Address,
	}

	if err := s.DB.Create(&supplier).Error; err != nil {
		return nil, err
	}

	return &supplier, nil
}

func (s *SupplierService) GetAll(page int, pageSize int) ([]models.Supplier, int64, error) {
	var (
		suppliers []models.Supplier
		total     int64
	)

	if err := s.DB.Model(&models.Supplier{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := s.DB.Limit(pageSize).Offset(offset).Find(&suppliers).Error; err != nil {
		return nil, 0, err
	}

	return suppliers, total, nil
}

func (s *SupplierService) Update(
	supplier *models.Supplier,
	req dto.UpdateSupplierRequest,
) error {

	if req.Name != nil {
		supplier.Name = *req.Name
	}
	if req.Email != nil {
		supplier.Email = *req.Email
	}
	if req.Address != nil {
		supplier.Address = *req.Address
	}

	if err := s.DB.Save(supplier).Error; err != nil {
		return err
	}

	return nil
}

func (s *SupplierService) Delete(supplier *models.Supplier) error {
	if err := s.DB.Delete(supplier).Error; err != nil {
		return err
	}
	return nil
}
