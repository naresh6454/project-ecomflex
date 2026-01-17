package route

import (
	"github.com/gin-gonic/gin"
	"github.com/naresh6454/ecomflex-backend/internal/api/handler"
	"github.com/naresh6454/ecomflex-backend/internal/api/middleware"
)

// ConfigureAuthRoutes sets up authentication routes
func ConfigureAuthRoutes(router *gin.RouterGroup, authHandler *handler.AuthHandler, authMiddleware *middleware.AuthMiddleware) {
	// Public routes
	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)
	router.POST("/refresh", authHandler.RefreshToken)
	router.POST("/google/login", authHandler.GoogleLogin)
	
	// Protected routes
	router.POST("/logout", authMiddleware.Authenticate(), authHandler.Logout)
}