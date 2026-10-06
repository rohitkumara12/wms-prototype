package services

import (
	"wms-prototype/models"
	"wms-prototype/views"

	"gorm.io/gorm"
)

func CreateOrder(request views.OrderRequest) error {

	return models.DB.Transaction(func(tx *gorm.DB) error {

		order := models.Order{
			OrderNumber: request.OrderNumber,
			Status:      "WAITING ..",
		}
		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		for _, item := range request.Items {
			orderItem := models.OrderItem{
				OrderID:   order.ID,
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
			}

			if err := tx.Save(&orderItem).Error; err != nil {
				return err
			}
		}
		return nil

	})
}

func GetOrders() ([]models.Order, error) {

	var orders []models.Order

	if err := models.DB.
		Preload("OrderItems").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
