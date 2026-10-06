package services

import (
	"wms-prototype/models"
	"wms-prototype/views"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(request views.RegisterRequest) error {

	hashedpassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	RegisterUser := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: string(hashedpassword),
		Role:     request.Role,
	}

	if err := models.DB.Create(&RegisterUser).Error; err != nil {
		return err
	}

	return nil
}
