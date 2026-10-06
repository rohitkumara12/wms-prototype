package models

import "gorm.io/gorm"

type WaveItem struct {
	gorm.Model

	WaveID      uint
	ProductID   uint
	RequiredQty uint
	PickedQty   uint

	Product Product `gorm:"foreignKey:ProductID"`
}
