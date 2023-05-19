package routes

import (
	"github.com/gin-gonic/gin"
	"goenv/controller"
	"goenv/middleware"
)

func SetupRoutes(router *gin.Engine) {

	router.POST("/login", controller.Login())
	router.POST("/refreshToken", controller.RefreshToken())
	v1 := router.Group("/api/v1")
	v1.Use(middleware.AuthorizeJWT())
	{
		v1.POST("/employee", controller.CreateEmployee())
		v1.GET("/employee/:id", controller.GetEmployee())
		v1.GET("/employees", controller.GetEmployees())
		v1.PUT("/employee/:id", controller.UpdateEmployee())
		v1.DELETE("/employee/:id", controller.DeleteEmployee())

	}

}
