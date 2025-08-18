package main

import (
	"github.com/gin-gonic/gin"
	"github.com/imbivek08/web-service/config"
	"github.com/imbivek08/web-service/routes"
)

func main() {
	router := gin.Default()

	config.ConnectDatabase()
	//public route
	routes.AuthRutes(router)
	//protected route
	routes.UserRoutes(router)
	routes.ProductRoute(router)
	routes.OrderRoutes(router)
	router.Run(":8000")
}
