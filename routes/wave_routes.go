package routes

import (
	"wms-prototype/controllers"

	"github.com/gin-gonic/gin"
)

func WaveRoutes(api *gin.RouterGroup) {

	api.POST("/waves", controllers.CreateWave)
	api.PUT("/waves/:id/status", controllers.UpdateWaveStatus)
	api.GET("/waves", controllers.GetWaves)
}
