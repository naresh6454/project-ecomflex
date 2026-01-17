package route

import (
	"github.com/gin-gonic/gin"
	"github.com/naresh6454/ecomflex-backend/internal/api/handler"
	"github.com/naresh6454/ecomflex-backend/internal/api/middleware"
)

// ConfigureAdminRoutes sets up admin routes
func ConfigureAdminRoutes(
	router *gin.RouterGroup,
	adminHandler *handler.AdminHandler,
	authMiddleware *middleware.AuthMiddleware,
) {
	// Admin routes (requires authentication and admin role)
	router.Use(authMiddleware.RequireRole("admin")) // Use admin role
	
	// Influencer management
	router.GET("/influencers", adminHandler.GetInfluencers)
	router.GET("/influencers/:id", adminHandler.GetInfluencer)
	router.PUT("/influencers/:id/status", adminHandler.UpdateInfluencerStatus)
}