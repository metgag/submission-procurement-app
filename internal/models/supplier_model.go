package models

import "gorm.io/gorm"

type Supplier struct {
	Name    string `gorm:"type:varchar(70);not null"`
	Email   string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Address string `gorm:"type:varchar(255);not null"`
	gorm.Model
}
