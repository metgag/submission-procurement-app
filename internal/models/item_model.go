package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name string `gorm:"varchar(100);not null"`
}
