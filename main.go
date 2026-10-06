package main

import (
	"wms-prototype/models"
	"wms-prototype/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	models.ConnectedDB()

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")
}
