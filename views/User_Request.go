package views

import "gorm.io/gorm"

type RegisterRequest struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Role     string `json:"role"`
}
