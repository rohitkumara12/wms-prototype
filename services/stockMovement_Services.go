package services

import (
	"errors"
	"wms-prototype/models"
	"wms-prototype/views"

	"gorm.io/gorm"
)

func StockIn(request views.StockMovementRequest, userID uint) error {

	return models.DB.Transaction(func(tx *gorm.DB) error {

		var product models.Product

		if err := tx.First(&product, request.ProductID).Error; err != nil {
			return errors.New("product not found")
		}

		stockbefore := product.Stock
		//tambah stok
		product.Stock += int(request.Quantity)

		if err := tx.Save(&product).Error; err != nil {
			return err
		}

		stockMovement := models.StockMovement{
			ProductID:   product.ID,
			Type:        "STOCK_IN",
			Quantity:    int(request.Quantity),
			StockBefore: stockbefore,
			StockAfter:  product.Stock,
			CreatedBy:   userID,
		}

		if err := tx.Create(&stockMovement).Error; err != nil {
			return err
		}

		return nil
	})
}

func StockOut(request views.StockMovementRequest, userID uint) error {
	return models.DB.Transaction(func(tx *gorm.DB) error {
		var product models.Product

		if err := tx.First(&product, request.ProductID).Error; err != nil {
			return errors.New("Product not found")
		}
		stockbefore := product.Stock

		if product.Stock < int(request.Quantity) {
			return errors.New("Insufficient Stock")
		}
		product.Stock -= int(request.Quantity)

		if err := tx.Save(&product).Error; err != nil {
			return err
		}

		stockmovement := models.StockMovement{
			ProductID:   product.ID,
			Type:        "STOCK_OUT",
			Quantity:    int(request.Quantity),
			StockBefore: stockbefore,
			StockAfter:  product.Stock,
			CreatedBy:   userID,
		}
		if err := tx.Create(&stockmovement).Error; err != nil {
			return err
		}
		return nil
	})
}
func GetStockMovement() ([]models.StockMovement, error) {
	var movements []models.StockMovement

	if err := models.DB.Order("created_at DESC").Find(&movements).Error; err != nil {
		return nil, err
	}
	return movements, nil
}
