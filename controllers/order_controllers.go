package controllers

import (
	"net/http"
	"wms-prototype/services"
	"wms-prototype/views"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var request views.OrderRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"data":   nil,
			"error":  []string{err.Error()},
		})
		return
	}

	//panggil service
	if err := services.CreateOrder(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"data":   nil,
			"error":  []string{err.Error()},
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": 201,
		"data":   "status created successfully",
		"error":  nil,
	})
}

func GetOrders(c *gin.Context) {
	var response views.BaseResponse

	products, err := services.GetOrders()
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Error = append(response.Error, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Status = http.StatusOK
	response.Data = products
	c.JSON(http.StatusOK, response)
}
