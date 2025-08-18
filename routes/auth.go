package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/imbivek08/web-service/controllers"
)

func AuthRutes(r *gin.Engine) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", controllers.Register)
		authGroup.POST("/login", controllers.Login)
	}
}
