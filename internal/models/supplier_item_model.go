package models

import "gorm.io/gorm"

type SupplierItem struct {
	gorm.Model
	SupplierID        uint  `gorm:"not null;index;uniqueIndex:uniq_supplier_item"`
	ItemID            uint  `gorm:"not null;index;uniqueIndex:uniq_supplier_item"`
	Price             int64 `gorm:"not null"`
	Stock             int   `gorm:"not null"`
	Supplier          Supplier
	Item              Item
	PurchasingDetails []PurchasingDetail `gorm:"foreignKey:SupplierItemID;constraint:OnDelete:CASCADE"`
}
