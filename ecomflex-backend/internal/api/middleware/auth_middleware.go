package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/naresh6454/ecomflex-backend/internal/api/response"
	"github.com/naresh6454/ecomflex-backend/internal/domain/entity"
	"github.com/naresh6454/ecomflex-backend/internal/domain/service"
	"github.com/google/uuid"
	loggerPkg "github.com/naresh6454/ecomflex-backend/internal/util/logger"
)

// AuthMiddleware handles authentication and authorization
type AuthMiddleware struct {
	authService service.AuthService
	logger      loggerPkg.Logger
}

// NewAuthMiddleware creates a new AuthMiddleware
func NewAuthMiddleware(authService service.AuthService, logger loggerPkg.Logger) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
		logger:      logger,
	}
}

// Authenticate authenticates a user using JWT token
func (m *AuthMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check for special superadmin authentication header
		superadminHeader := c.GetHeader("X-Superadmin-Auth")
		if superadminHeader == "ecomflex-superadmin-direct-auth" {
			// This is a superadmin request, create a mock admin user
			superadminID, _ := uuid.Parse("00000000-0000-0000-0000-000000000000") // Use a fixed UUID for superadmin
			tenantID, _ := uuid.Parse("00000000-0000-0000-0000-000000000001")     // Use a fixed UUID for tenant

			// Create a superadmin user according to your entity model
			// Adjust fields as necessary to match your User entity structure
			superadminUser := &entity.User{
				ID:       superadminID,
				Email:    "admin@ecomflex.com",
				FullName: "System Administrator",
				Role:     "admin", // Adjust according to your role type
				TenantID: tenantID,
			}

			// Set user in context
			c.Set("user", superadminUser)
			c.Set("userID", superadminUser.ID.String())
			c.Set("email", superadminUser.Email)
			c.Set("role", string(superadminUser.Role))
			c.Set("tenantID", superadminUser.TenantID.String())
			c.Set("tokenID", "superadmin-token-id")
			c.Set("isSuperAdmin", true)

			c.Next()
			return
		}

		// Standard JWT authentication
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Error(c, http.StatusUnauthorized, "Authorization header is required", nil)
			c.Abort()
			return
		}
		
		// Parse token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Error(c, http.StatusUnauthorized, "Invalid authorization format", nil)
			c.Abort()
			return
		}
		
		token := parts[1]
		
		// Validate token and get user
		user, err := m.authService.GetUserFromToken(c.Request.Context(), token)
		if err != nil {
			m.logger.Error("Failed to validate token", err)
			response.Error(c, http.StatusUnauthorized, "Invalid token", nil)
			c.Abort()
			return
		}
		
		// Check if user is active
		if !user.IsActive() {
			response.Error(c, http.StatusForbidden, "Your account is inactive", nil)
			c.Abort()
			return
		}
		
		// Set user in context
		c.Set("user", user)
		c.Set("userID", user.ID.String())
		c.Set("email", user.Email)
		c.Set("role", string(user.Role))
		c.Set("tenantID", user.TenantID.String())
		
		// Extract tokenID from claims and set in context
		c.Set("tokenID", "token-id-placeholder")
		
		c.Next()
	}
}

// RequireRole checks if the user has the required role
func (m *AuthMiddleware) RequireRole(roles ...entity.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user from context
		user, exists := c.Get("user")
		if !exists {
			response.Error(c, http.StatusUnauthorized, "Unauthorized", nil)
			c.Abort()
			return
		}
		
		// Check if user has one of the required roles
		userObj := user.(*entity.User)
		hasRole := false
		
		for _, role := range roles {
			if userObj.Role == role {
				hasRole = true
				break
			}
		}
		
		if !hasRole {
			response.Error(c, http.StatusForbidden, "Insufficient permissions", nil)
			c.Abort()
			return
		}
		
		c.Next()
	}
}