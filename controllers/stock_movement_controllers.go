package controllers

import (
	"net/http"
	"wms-prototype/services"
	"wms-prototype/views"

	"github.com/gin-gonic/gin"
)

func StockIn(c *gin.Context) {
	var request views.StockMovementRequest
	var response views.BaseResponse

	if err := c.ShouldBindJSON(&request); err != nil {
		response.Status = http.StatusBadRequest
		response.Error = append(response.Error, err.Error())

		c.JSON(http.StatusBadRequest, response)
		return
	}

	userID := c.MustGet("userID").(uint)

	if err := services.StockIn(request, userID); err != nil {
		response.Status = http.StatusBadRequest
		response.Error = append(response.Error, err.Error())

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response.Status = http.StatusOK
	response.Data = "Stock In Success"

	c.JSON(http.StatusOK, response)
}

func StockOut(c *gin.Context) {
	var request views.StockMovementRequest
	var response views.BaseResponse

	if err := c.ShouldBindJSON(&request); err != nil {
		response.Status = http.StatusBadRequest
		response.Error = append(response.Error, err.Error())
		c.JSON(http.StatusBadRequest, response)
	}

	userID := c.MustGet("userID").(uint)

	if err := services.StockOut(request, userID); err != nil {
		response.Status = http.StatusBadRequest
		response.Error = append(response.Error, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response.Status = http.StatusOK
	response.Data = "Stock out Success"

	c.JSON(http.StatusOK, response)

}

func GetStockMovement(c *gin.Context) {
	var response views.BaseResponse

	movements, err := services.GetStockMovement()
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Error = append(response.Error, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Status = http.StatusOK
	response.Data = movements
	c.JSON(http.StatusOK, movements)
}
