package models

import "gorm.io/gorm"

type Wave struct {
	gorm.Model

	WaveNumber string `gorm:"unique"`
	Status     string
	WaveItems  []WaveItem
	WaveOrders []WaveOrder
}
