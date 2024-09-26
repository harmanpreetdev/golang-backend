package routes

import (
	"demo/controllers"
	"demo/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// User routes Group
	api := router.Group("/api/v1")

	api.POST("/register", controllers.RegisterUser)
	api.POST("/login", controllers.LoginUser)

	api.Use(middlewares.JWTMiddleware())
	{
		api.GET("/users", controllers.GetUsers)
		api.POST("/users", controllers.CreateUser)
	}

}
