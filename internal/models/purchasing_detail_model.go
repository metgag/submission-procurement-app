package models

import "gorm.io/gorm"

type PurchasingDetail struct {
	gorm.Model
	PurchasingID   uint  `gorm:"index;not null"`
	SupplierItemID uint  `gorm:"index;not null"`
	Quantity       int   `gorm:"not null"`
	Subtotal       int64 `gorm:"not null"`
	Purchasing     Purchasing
	SupplierItem   SupplierItem
}
