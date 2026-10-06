package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {

	api := r.Group("/api")

	ProductRoutes(api)
	UserRoutes(api)
	StockMovementRoutes(api)
	OrderRoutes(api)
	WaveRoutes(api)

}
