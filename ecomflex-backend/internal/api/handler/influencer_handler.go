package handler

import (
	"net/http"
	"strconv"
	
	"github.com/gin-gonic/gin"
	"github.com/naresh6454/ecomflex-backend/internal/domain/service"
	"github.com/naresh6454/ecomflex-backend/internal/util/logger"
)

// InfluencerHandler handles influencer-related requests
type InfluencerHandler struct {
	influencerService service.InfluencerService
	logger            logger.Logger
}

// NewInfluencerHandler creates a new InfluencerHandler
func NewInfluencerHandler(influencerService service.InfluencerService, logger logger.Logger) *InfluencerHandler {
	return &InfluencerHandler{
		influencerService: influencerService,
		logger:            logger,
	}
}

// GetReferralCode returns the influencer's referral code
func (h *InfluencerHandler) GetReferralCode(c *gin.Context) {
	userID := c.GetString("userID")
	tenantID := c.GetString("tenantID")
	if userID == "" || tenantID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	
	referralCode, err := h.influencerService.GetReferralCode(c, userID, tenantID)
	if err != nil {
		h.logger.Error("Failed to get referral code", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get referral code"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"referral_code": referralCode,
		},
	})
}

// GetReferralStats returns the influencer's referral statistics
func (h *InfluencerHandler) GetReferralStats(c *gin.Context) {
	userID := c.GetString("userID")
	tenantID := c.GetString("tenantID")
	if userID == "" || tenantID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	
	stats, err := h.influencerService.GetReferralStats(c, userID, tenantID)
	if err != nil {
		h.logger.Error("Failed to get referral stats", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get referral stats"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": stats,
	})
}

// GetRecentReferrals returns the influencer's recent referrals
func (h *InfluencerHandler) GetRecentReferrals(c *gin.Context) {
	userID := c.GetString("userID")
	tenantID := c.GetString("tenantID")
	if userID == "" || tenantID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	
	limitStr := c.DefaultQuery("limit", "5")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 5
	}
	
	referrals, err := h.influencerService.GetRecentReferrals(c, userID, tenantID, limit)
	if err != nil {
		h.logger.Error("Failed to get recent referrals", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get recent referrals"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"referrals": referrals,
		},
	})
}

// GetDashboardData returns all data needed for the influencer dashboard
func (h *InfluencerHandler) GetDashboardData(c *gin.Context) {
	userID := c.GetString("userID")
	tenantID := c.GetString("tenantID")
	if userID == "" || tenantID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	
	h.logger.Info("Getting dashboard data", map[string]interface{}{
		"userID": userID,
		"tenantID": tenantID,
	})
	
	dashboardData, err := h.influencerService.GetDashboardData(c, userID, tenantID)
	if err != nil {
		h.logger.Error("Failed to get dashboard data", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get dashboard data"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": dashboardData,
	})
}