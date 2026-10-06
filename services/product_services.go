package services

import (
	"wms-prototype/models"
	"wms-prototype/views"
)

func CreateProduct(request views.ProductRequest) error {

	product := models.Product{
		SKU:      request.SKU,
		Name:     request.Name,
		Stock:    request.Stock,
		Location: request.Location,
	}

	if err := models.DB.Create(&product).Error; err != nil {
		return err
	}
	return nil
}

func GetProducts() ([]models.Product, error) {

	var products []models.Product

	if err := models.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
func GetProductbyID(id uint) (models.Product, error) {
	var product models.Product

	if err := models.DB.First(&product, id).Error; err != nil {
		return product, err
	}
	return product, nil
}
func UpdateProduct(id uint, request views.ProductRequest) error {

	var product models.Product

	if err := models.DB.First(&product, id).Error; err != nil {
		return err
	}
	product.SKU = request.SKU
	product.Name = request.Name
	product.Stock = request.Stock
	product.Location = product.Location

	if err := models.DB.Save(&product).Error; err != nil {
		return err
	}
	return nil
}

func DeleteProduct(id uint) error {

	var product models.Product

	if err := models.DB.First(&product, id).Error; err != nil {
		return err
	}

	if err := models.DB.Delete(&product).Error; err != nil {
		return err
	}

	return nil
}
