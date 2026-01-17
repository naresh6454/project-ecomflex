package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/naresh6454/ecomflex-backend/internal/api/response"
	"github.com/naresh6454/ecomflex-backend/internal/domain/entity"
	"github.com/naresh6454/ecomflex-backend/internal/domain/service"
	loggerPkg "github.com/naresh6454/ecomflex-backend/internal/util/logger"
)

// AdminHandler handles admin-related requests
type AdminHandler struct {
	userService service.UserService
	logger      loggerPkg.Logger
}

// NewAdminHandler creates a new AdminHandler
func NewAdminHandler(userService service.UserService, logger loggerPkg.Logger) *AdminHandler {
	return &AdminHandler{
		userService: userService,
		logger:      logger,
	}
}

// GetInfluencers lists influencers with pagination and filters
func (h *AdminHandler) GetInfluencers(c *gin.Context) {
	// Parse query parameters
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")
	status := c.DefaultQuery("status", "pending_approval")
	// We're not using these yet but keeping them for future implementation
	// search := c.DefaultQuery("search", "")
	// fromDate := c.DefaultQuery("from_date", "")

	// Convert string parameters to int
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	// Calculate offset
	offset := (page - 1) * pageSize

	var influencers []*entity.User
	var total int

	// Get influencers from service based on status
	if status == "pending_approval" {
		influencers, total, err = h.userService.GetPendingInfluencers(
			c.Request.Context(), 
			pageSize, 
			offset,
		)
	} else {
		// For now, we only handle pending_approval status directly
		// This can be expanded when more methods are added to UserService
		h.logger.Error("Unsupported status filter", nil)
		response.Error(c, http.StatusBadRequest, "Unsupported status filter. Only 'pending_approval' is supported currently.", nil)
		return
	}

	if err != nil {
		h.logger.Error("Failed to get influencers", err)
		response.Error(c, http.StatusInternalServerError, "Failed to get influencers", err)
		return
	}

	// Calculate total pages
	totalPages := (total + pageSize - 1) / pageSize

	// Create response
	resp := map[string]interface{}{
		"data":        influencers,
		"total":       total,
		"page":        page,
		"page_size":   pageSize,
		"total_pages": totalPages,
	}

	// Return success response
	response.Success(c, http.StatusOK, "Influencers retrieved successfully", resp)
}

// UpdateInfluencerStatus updates the status of an influencer
func (h *AdminHandler) UpdateInfluencerStatus(c *gin.Context) {
	// Parse influencer ID from URL
	influencerIDStr := c.Param("id")
	influencerID, err := uuid.Parse(influencerIDStr)
	if err != nil {
		h.logger.Error("Invalid influencer ID", err)
		response.Error(c, http.StatusBadRequest, "Invalid influencer ID", err)
		return
	}

	// Parse request body
	var req service.InfluencerStatusUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Failed to bind request", err)
		response.Error(c, http.StatusBadRequest, "Invalid request", err)
		return
	}

	// Update influencer status
	err = h.userService.UpdateInfluencerStatus(c.Request.Context(), influencerID, req)
	if err != nil {
		h.logger.Error("Failed to update influencer status", err)
		response.Error(c, http.StatusInternalServerError, "Failed to update influencer status", err)
		return
	}

	// Return success response
	response.Success(c, http.StatusOK, "Influencer status updated successfully", nil)
}

// GetInfluencer gets an influencer by ID
func (h *AdminHandler) GetInfluencer(c *gin.Context) {
	// Parse influencer ID from URL
	influencerIDStr := c.Param("id")
	influencerID, err := uuid.Parse(influencerIDStr)
	if err != nil {
		h.logger.Error("Invalid influencer ID", err)
		response.Error(c, http.StatusBadRequest, "Invalid influencer ID", err)
		return
	}

	// Get influencer from service
	influencer, err := h.userService.GetUserByID(c.Request.Context(), influencerID)
	if err != nil {
		h.logger.Error("Failed to get influencer", err)
		response.Error(c, http.StatusInternalServerError, "Failed to get influencer", err)
		return
	}

	// Return success response
	response.Success(c, http.StatusOK, "Influencer retrieved successfully", influencer)
}