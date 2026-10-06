package models

import "gorm.io/gorm"

type StockMovement struct {
	gorm.Model
	ProductID   uint `gorm:"not null"`
	Product     Product
	Type        string `gorm:"not null"`
	Quantity    int    `gorm:"not null"`
	StockBefore int
	StockAfter  int
	CreatedBy   uint
}
