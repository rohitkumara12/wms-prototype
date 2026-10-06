package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	OrderNumber string `gorm:"unique;not null"`
	Status      string `gorm:"default:'Waiting'"`

	OrderItems []OrderItem `gorm:"foreignKey:OrderID"`
	Waveorders []WaveOrder
}
