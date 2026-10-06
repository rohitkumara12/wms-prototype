package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var Secretkeys = []byte("wms-prototype-secret-key")

func GenerateJWT(userID uint, email string, role string) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(Secretkeys)
}
