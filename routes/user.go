package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/imbivek08/web-service/controllers"
	"github.com/imbivek08/web-service/middlewares"
)

func UserRoutes(r *gin.Engine) {
	userGroup := r.Group("/user")
	userGroup.Use(middlewares.AuthMiddleware())
	{
		userGroup.PUT("/update", controllers.UpdateUser)
		userGroup.GET("/profile", controllers.GetUser)
		userGroup.DELETE("/delete", controllers.DeleteUser)
	}
}
