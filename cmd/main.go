package main

import (
	"github.com/gin-gonic/gin"
	"github.com/imbivek08/web-service/config"
	"github.com/imbivek08/web-service/models"
	"github.com/imbivek08/web-service/routes"
)

func main() {
	router := gin.Default()

	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.User{})
	routes.AuthRutes(router)
	router.Run(":8080")
}
