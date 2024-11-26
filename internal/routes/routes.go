package route

import (
	"github.com/gin-gonic/gin"
	"saaster.tech/crud/internal/handler"
)

func RegisterRoutes(router *gin.Engine, userHandler *handler.UserHandler) {
	api := router.Group("/api/v1")
	{
		api.GET("/users", userHandler.GetUsers)
		api.POST("/users", userHandler.AddUser)
	}
}
