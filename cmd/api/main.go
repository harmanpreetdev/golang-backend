package main

import (
	"demo/config"
	"demo/database"
	"demo/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	database.ConnectDatabase()

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	routes.SetupRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(router.Run(":" + port))
}
