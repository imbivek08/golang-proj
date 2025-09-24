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
	config.InitRedis()
	//for development using automigrate for quick setup
	config.DB.AutoMigrate(models.User{})
	config.DB.AutoMigrate(models.Product{})
	config.DB.AutoMigrate(models.Order{})
	config.DB.AutoMigrate(models.OrderItem{})
	//public route
	routes.AuthRutes(router)
	//protected route
	routes.UserRoutes(router)
	routes.ProductRoute(router)
	routes.OrderRoutes(router)
	router.Run(":8000")
}
