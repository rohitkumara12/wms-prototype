package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	SKU      string `gorm:"unique;not null"`
	Name     string `gorm:"not null"`
	Stock    int    `gorm:"default:0"`
	Location string
}
