package services

import (
	"wms-prototype/models"
	"wms-prototype/utils"
	"wms-prototype/views"

	"golang.org/x/crypto/bcrypt"
)

func Login(request views.LoginRequest) (string, error) {
	var user models.User

	//cari berdasarkan email
	if err := models.DB.Where("email = ?", request.Email).First(&user).Error; err != nil {
		return "", err
	}
	//cek/bandingkan password
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(request.Password),
	); err != nil {
		return "", err
	}
	//GenerateJWT
	token, err := utils.GenerateJWT(user.ID, user.Email, user.Role)
	if err != nil {
		return "", err
	}
	return token, nil
}
