package services

import (
	"gorm.io/gorm"

	"github.com/metgag/procurement-api-example/internal/dto"
	"github.com/metgag/procurement-api-example/internal/models"
)

type SupplierItemService struct {
	DB *gorm.DB
}

func NewSupplierItemService(db *gorm.DB) *SupplierItemService {
	return &SupplierItemService{DB: db}
}

// Create assigns an item to a supplier
func (s *SupplierItemService) Create(req dto.CreateSupplierItemRequest) (*models.SupplierItem, error) {
	supplierItem := models.SupplierItem{
		SupplierID: req.SupplierID,
		ItemID:     req.ItemID,
		Price:      req.Price,
		Stock:      req.Stock,
	}

	if err := s.DB.Create(&supplierItem).Error; err != nil {
		return nil, err
	}

	return &supplierItem, nil
}

// GetAll returns supplier items, optionally filtered by supplierID
func (s *SupplierItemService) GetAll(supplierID, page, pageSize int) ([]models.SupplierItem, int64, error) {
	var (
		items []models.SupplierItem
		total int64
	)

	query := s.DB.Model(&models.SupplierItem{}).Preload("Item").Preload("Supplier")

	if supplierID > 0 {
		query = query.Where("supplier_id = ?", supplierID)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := query.Limit(pageSize).Offset(offset).Find(&items).Error; err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

// Update modifies price or stock of a supplier item
func (s *SupplierItemService) Update(supplierItem *models.SupplierItem, req dto.UpdateSupplierItemRequest) error {
	if req.Price != nil {
		supplierItem.Price = *req.Price
	}
	if req.Stock != nil {
		supplierItem.Stock = *req.Stock
	}

	if err := s.DB.Save(supplierItem).Error; err != nil {
		return err
	}

	return nil
}

// Delete removes a supplier item
func (s *SupplierItemService) Delete(supplierItem *models.SupplierItem) error {
	if err := s.DB.Delete(supplierItem).Error; err != nil {
		return err
	}
	return nil
}
