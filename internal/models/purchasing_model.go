package models

import (
	"time"

	"gorm.io/gorm"
)

type Purchasing struct {
	gorm.Model
	OrderDate  time.Time `gorm:"not null"`
	SupplierID uint      `gorm:"index;not null"`
	UserID     uint      `gorm:"index;not null"`
	GrandTotal int64     `gorm:"not null"`
	Supplier   Supplier
	User       User
	Details    []PurchasingDetail `gorm:"constraint:OnDelete:CASCADE"`
}
