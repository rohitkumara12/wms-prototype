package routes

import (
	"wms-prototype/controllers"
	"wms-prototype/middleware"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(api *gin.RouterGroup) {

	product := api.Group("/products")

	product.Use(middleware.AuthMiddleware())

	product.POST("", controllers.CreateProduct)
	product.GET("", controllers.GetProducts)
	product.GET("/:id", controllers.GetProductbyID)
	product.PUT("/:id", controllers.UpdateProduct)
	product.DELETE("/:id", controllers.DeleteProduct)
}
