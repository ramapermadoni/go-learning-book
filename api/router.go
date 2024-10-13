package api

import (
	handlers "go-learning-book/api/handler"
	"go-learning-book/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// User routes
	userGroup := r.Group("/user")
	{
		userGroup.POST("/login", handlers.Login)                // Login
		userGroup.POST("/refresh-token", handlers.RefreshToken) // Refresh Token
	}

	// Protected routes
	protectedGroup := r.Group("/protected")
	protectedGroup.Use(middlewares.JwtMiddleware())
	{
		// Add protected routes here
	}

	return r
}
