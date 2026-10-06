package models

import "gorm.io/gorm"

type WaveOrder struct {
	gorm.Model

	WaveID  uint `gorm:"not null"`
	OrderID uint `gorm:"not null"`

	Wave  Wave  `gorm:"foreignkey:WaveID"`
	Order Order `gorm:"foreignkey:OrderID"`
}
