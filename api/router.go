package api

import (
	"auth-service/api/handlers"
	"github.com/gin-gonic/gin"
)

func NewRouter(userHandler *handlers.UserHandler) *gin.Engine {
	router := gin.Default()

	// Middleware
	router.Use(handlers.AuthMiddleware())

	v1 := router.Group("/v1")
	{
		v1.POST("/register/user", userHandler.Register)
		v1.POST("/login", userHandler.Login)
		v1.GET("/user", userHandler.GetUser)
		v1.GET("/users", userHandler.GetUsers)
	}

	return router
}
