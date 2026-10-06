package controllers

import (
	"net/http"
	"strconv"
	"wms-prototype/services"
	"wms-prototype/views"

	"github.com/gin-gonic/gin"
)

func CreateWave(c *gin.Context) {
	var request views.WaveRequest
	var response views.BaseResponse

	if err := c.ShouldBindJSON(&request); err != nil {
		response.Status = http.StatusBadRequest
		response.Error = append(response.Error, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	wave, err := services.CreateWaves(request)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Error = append(response.Error, err.Error())
		c.JSON(http.StatusBadRequest, response)
	}

	response.Status = http.StatusOK
	response.Data = wave
	c.JSON(http.StatusOK, response)
}

func UpdateWaveStatus(c *gin.Context) {
	var request views.WaveStatusRequest
	var response views.BaseResponse

	waveID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Error = append(response.Error, "invalid wave id")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//ambil data status dari JSON
	if err := c.ShouldBindJSON(&request); err != nil {
		response.Status = http.StatusBadRequest
		response.Error = append(response.Error, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//panggil service

	if err := services.UpdateStatusWave(uint(waveID), request); err != nil {
		response.Status = http.StatusBadRequest
		response.Error = append(response.Error, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response.Status = http.StatusOK
	response.Data = "Wave Status Updated Successfully"
	c.JSON(http.StatusOK, response)
}

func GetWaves(c *gin.Context) {
	var response views.BaseResponse

	waves, err := services.GetWaves()

	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Error = append(response.Error, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Status = http.StatusOK
	response.Data = waves
	c.JSON(http.StatusOK, response)
}
