package main

import (
	"goenv/configs"

	"github.com/gin-gonic/gin"
	"goenv/routes"
)

func main() {
	router := gin.Default()

	configs.ConnectDB()
	routes.SetupRoutes(router)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
