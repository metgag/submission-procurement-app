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
	if len(req.Items) == 0 {
		return nil, nil, errors.New("purchasing must have at least one item")
	}

	var (
		purchase  models.Purchasing
		itemsRes  []dto.PurchasingItemResponse
		tempItems []dto.PurchasingItemResponse
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
				return errors.New("insufficient item stock")
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

			// langsung return err, cek apakah ter rollback atau berhasil
			// prior medium
			if err := tx.
				Model(&supplierItem).
				Where("id = ?", supplierItem.ID).
				UpdateColumn("stock", gorm.Expr("stock - ?", item.Quantity)).
				// Update("stock", supplierItem.Stock-item.Quantity).
				Error; err != nil {
				return err
			}

			grandTotal += subtotal
			tempItems = append(tempItems, dto.PurchasingItemResponse{
				ItemID:   supplierItem.ItemID,
				Quantity: item.Quantity,
				Subtotal: subtotal,
			})
		}

		// cek juga untuk commit nya apa berhasil
		return tx.Model(&purchase).Update("grand_total", grandTotal).Error
	})
	if err != nil {
		return nil, nil, err
	}

	itemsRes = tempItems
	return &purchase, itemsRes, nil
}
