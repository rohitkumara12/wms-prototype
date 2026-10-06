package controllers

import (
	"net/http"
	"wms-prototype/services"
	"wms-prototype/views"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	var request views.LoginRequest
	var response views.BaseResponse

	if err := c.ShouldBindJSON(&request); err != nil {
		response.Status = http.StatusBadRequest
		response.Error = append(response.Error, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//panggil service login

	token, err := services.Login(request)
	if err != nil {
		response.Status = http.StatusUnauthorized
		response.Error = append(response.Error, "Email atau Password salah")
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	//login berhasil
	response.Status = http.StatusOK
	response.Data = gin.H{
		"token": token,
	}

	c.JSON(http.StatusOK, response)
}
