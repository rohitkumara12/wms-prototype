package routes

import (
	"wms-prototype/controllers"
	"wms-prototype/middleware"

	"github.com/gin-gonic/gin"
)

func StockMovementRoutes(api *gin.RouterGroup) {

	api.Use(middleware.AuthMiddleware())

	api.POST("/stock-in", controllers.StockIn)
	api.POST("/stock-out", controllers.StockOut)
	api.GET("/stock-movements", controllers.GetStockMovement)
}
