// File: /ecomflex-backend/internal/api/route/influencer_routes.go
package route

import (
	"github.com/gin-gonic/gin"
	"github.com/naresh6454/ecomflex-backend/internal/api/handler"
)

// RegisterInfluencerRoutes sets up all routes for influencer features
func RegisterInfluencerRoutes(router *gin.RouterGroup, influencerHandler *handler.InfluencerHandler, authMiddleware gin.HandlerFunc) {
	// Create a group for all influencer routes
	influencerGroup := router.Group("/influencer")
	
	// Apply authentication middleware
	influencerGroup.Use(authMiddleware)
	
	// Apply role checking middleware
	influencerGroup.Use(func(c *gin.Context) {
		// Get user role from context (should be set by auth middleware)
		role, exists := c.Get("userRole")
		if !exists || role != "influencer" {
			c.JSON(403, gin.H{"error": "Access denied: Influencer role required"})
			c.Abort()
			return
		}
		c.Next()
	})
	
	// Define API routes
	influencerGroup.GET("/dashboard", influencerHandler.GetDashboardData)
	influencerGroup.GET("/referral-code", influencerHandler.GetReferralCode)
	influencerGroup.GET("/stats", influencerHandler.GetReferralStats)
	influencerGroup.GET("/recent-referrals", influencerHandler.GetRecentReferrals)
}