package routes

import (
	"wms-prototype/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(api *gin.RouterGroup) {
	api.POST("/orders", controllers.CreateOrder)
	api.GET("/orders", controllers.GetOrders)
}
