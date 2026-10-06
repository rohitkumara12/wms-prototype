package controllers

import (
	"net/http"
	"wms-prototype/services"
	"wms-prototype/views"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {

	var Request views.RegisterRequest
	var response views.BaseResponse

	err := c.ShouldBindJSON(&Request)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Error = append(response.Error, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if err := services.RegisterUser(Request); err != nil {
		response.Status = http.StatusBadRequest
		response.Error = append(response.Error, "Failed To Register User")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	
	response.Status = http.StatusCreated
	response.Data = response
	c.JSON(http.StatusOK, response)
}
