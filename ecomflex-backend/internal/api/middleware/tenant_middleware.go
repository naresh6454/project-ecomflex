package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strings"
	loggerPkg "github.com/naresh6454/ecomflex-backend/internal/util/logger"
)

// TenantMiddleware handles tenant identification
type TenantMiddleware struct {
	logger loggerPkg.Logger
}

// NewTenantMiddleware creates a new TenantMiddleware
func NewTenantMiddleware(logger loggerPkg.Logger) *TenantMiddleware {
	return &TenantMiddleware{
		logger: logger,
	}
}

// ExtractTenantID extracts tenant ID from various sources
func (m *TenantMiddleware) ExtractTenantID() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tenantID string
		
		// Try to extract tenant ID from header
		tenantID = c.GetHeader("X-Tenant-ID")
		
		// If not found, check if we have a user in context (set by auth middleware)
		if tenantID == "" {
			if user, exists := c.Get("user"); exists {
				// Get tenant ID from authenticated user
				if userMap, ok := user.(map[string]interface{}); ok {
					if tid, exists := userMap["tenant_id"].(string); exists {
						tenantID = tid
					}
				}
			}
		}
		
		// If not found, try to extract from subdomain
		if tenantID == "" {
			host := c.Request.Host
			// Extract subdomain if present
			if strings.Contains(host, ".") {
				parts := strings.Split(host, ".")
				if len(parts) > 2 {
					subdomain := parts[0]
					
					// Skip common subdomains
					if subdomain != "www" && subdomain != "api" {
						// TODO: Look up tenant by subdomain from database
						// For now, we'll just use the default tenant ID
						// This should be replaced with actual lookup in production
					}
				}
			}
			
			// If still no tenant ID found, use default for public routes
			if tenantID == "" {
				tenantID = "00000000-0000-0000-0000-000000000000"
			}
		}
		
		// Validate UUID format
		_, err := uuid.Parse(tenantID)
		if err != nil {
			// Create a proper error message for logging
			logErr := fmt.Errorf("invalid tenant ID format %s: %w", tenantID, err)
			m.logger.Error("Invalid tenant ID format", logErr)
			
			// Fall back to default tenant ID
			tenantID = "00000000-0000-0000-0000-000000000000"
		}
		
		// Set tenant ID in context
		c.Set("tenantID", tenantID)
		
		c.Next()
	}
}