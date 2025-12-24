package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name          string         `gorm:"varchar(100);not null"`
	SupplierItems []SupplierItem `gorm:"foreignKey:ItemID;constraint:OnDelete:CASCADE"`
}
