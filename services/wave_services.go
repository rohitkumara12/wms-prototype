package services

import (
	"fmt"
	"wms-prototype/models"
	"wms-prototype/views"

	"gorm.io/gorm"
)

func CreateWaves(request views.WaveRequest) (*views.WaveResponse, error) {
	var response views.WaveResponse

	err := models.DB.Transaction(func(tx *gorm.DB) error {

		//buat wave
		wave := models.Wave{
			WaveNumber: request.WaveNumber,
			Status:     "WAITING",
		}
		if err := tx.Create(&wave).Error; err != nil {
			return err
		}
		var ordernumbers []string
		//loop semua order
		for _, orderID := range request.OrderIDs {
			var order models.Order

			if err := tx.First(&order, orderID).Error; err != nil {
				return err
			}

			//simpan relasi wave => order
			waveOrder := models.WaveOrder{
				WaveID:  wave.ID,
				OrderID: order.ID,
			}
			if err := tx.Create(&waveOrder).Error; err != nil {
				return err
			}

			//ambil item dari order
			var orderItems []models.OrderItem

			if err := tx.Where("order_id = ?", orderID).Find(&orderItems).Error; err != nil {
				return err
			}
			ordernumbers = append(ordernumbers, order.OrderNumber)

			//buat wave item

			for _, item := range orderItems {
				waveItem := models.WaveItem{
					WaveID:      wave.ID,
					ProductID:   item.ProductID,
					RequiredQty: item.Quantity,
					PickedQty:   0,
				}

				if err := tx.Create(&waveItem).Error; err != nil {
					return err
				}
			}
		}
		response.WaveNumber = wave.WaveNumber
		response.Status = wave.Status
		response.Orders = ordernumbers

		return nil
	})
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func UpdateStatusWave(waveID uint, request views.WaveStatusRequest) error {
	var wave models.Wave

	//cari wave berdasarkan ID
	if err := models.DB.First(&wave, waveID).Error; err != nil {
		return err
	}

	//cek status yang diperbolehkan
	if !IsValidWaveStatus(request.Status) {
		return fmt.Errorf("invalid wave status")
	}

	//update status
	wave.Status = request.Status

	if err := models.DB.Save(&wave).Error; err != nil {
		return err
	}
	return nil
}

func GetWaves() ([]models.Wave, error) {

	var waves []models.Wave

	if err := models.DB.Preload("WaveOrders.Order").
		Find(&waves).Error; err != nil {
		return nil, err
	}
	return waves, nil
}
func IsValidWaveStatus(status string) bool {
	return status == "WAITING" ||
		status == "PROCESSING" ||
		status == "COMPLETED"
}
