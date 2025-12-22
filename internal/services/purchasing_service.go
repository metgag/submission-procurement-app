package services

import (
	"errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/metgag/procurement-api-example/internal/dto"
	"github.com/metgag/procurement-api-example/internal/models"
)

type PurchasingService struct {
	DB *gorm.DB
}

func NewPurchasingService(db *gorm.DB) *PurchasingService {
	return &PurchasingService{DB: db}
}

func (s *PurchasingService) Create(
	userID uint,
	req dto.CreatePurchasingRequest,
) (*models.Purchasing, []dto.PurchasingItemResponse, error) {
	var (
		purchase models.Purchasing
		itemsRes []dto.PurchasingItemResponse
	)

	err := s.DB.Transaction(func(tx *gorm.DB) error {
		purchase = models.Purchasing{
			UserID:     userID,
			SupplierID: req.SupplierID,
			OrderDate:  time.Now(),
		}

		if err := tx.Create(&purchase).Error; err != nil {
			return err
		}

		var grandTotal int64
		for _, item := range req.Items {
			var supplierItem models.SupplierItem

			if err := tx.
				Clauses(clause.Locking{Strength: "UPDATE"}).
				Where("id = ? AND supplier_id = ?", item.SupplierItemID, req.SupplierID).
				First(&supplierItem).
				Error; err != nil {
				return err
			}

			if supplierItem.Stock < item.Quantity {
				return errors.New("unsufficient item stock")
			}

			subtotal := supplierItem.Price * int64(item.Quantity)
			detail := models.PurchasingDetail{
				PurchasingID:   purchase.ID,
				SupplierItemID: supplierItem.ID,
				Quantity:       item.Quantity,
				Subtotal:       subtotal,
			}

			if err := tx.Create(&detail).Error; err != nil {
				return err
			}

			if err := tx.
				Model(&supplierItem).
				Update("stock", supplierItem.Stock-item.Quantity).
				Error; err != nil {
				return err
			}

			grandTotal += subtotal
			itemsRes = append(itemsRes, dto.PurchasingItemResponse{
				ItemID:   supplierItem.ItemID,
				Quantity: item.Quantity,
				Subtotal: subtotal,
			})
		}

		return tx.Model(&purchase).Update("grand_total", grandTotal).Error
	})
	if err != nil {
		return nil, nil, err
	}

	return &purchase, itemsRes, nil
}
