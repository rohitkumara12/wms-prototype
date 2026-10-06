package routes

import (
	"wms-prototype/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(api *gin.RouterGroup) {

	api.POST("/register", controllers.RegisterUser)
	api.POST("/login", controllers.Login)
}
