package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/naresh6454/ecomflex-backend/internal/api/request"
	"github.com/naresh6454/ecomflex-backend/internal/api/response"
	"github.com/naresh6454/ecomflex-backend/internal/domain/service"
	loggerPkg "github.com/naresh6454/ecomflex-backend/internal/util/logger"
)

// AuthHandler handles authentication-related requests
type AuthHandler struct {
	authService service.AuthService
	logger      loggerPkg.Logger
}

// NewAuthHandler creates a new AuthHandler
func NewAuthHandler(authService service.AuthService, logger loggerPkg.Logger) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		logger:      logger,
	}
}

// Register handles user registration
func (h *AuthHandler) Register(c *gin.Context) {
	var req request.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Failed to bind register request", err)
		response.Error(c, http.StatusBadRequest, "Invalid request", err)
		return
	}
	
	// Get tenant ID from context (set by middleware)
	tenantIDStr := c.GetString("tenantID")
	if tenantIDStr == "" {
		// If not set, use a default tenant ID
		tenantIDStr = "00000000-0000-0000-0000-000000000000"
	}
	
	tenantID, err := uuid.Parse(tenantIDStr)
	if err != nil {
		h.logger.Error("Invalid tenant ID", err)
		response.Error(c, http.StatusBadRequest, "Invalid tenant ID", err)
		return
	}
	
	// Convert request to service request
	serviceReq := service.RegisterRequest{
		Email:         req.Email,
		Password:      req.Password,
		FullName:      req.FullName,
		Phone:         req.Phone,
		Role:          req.Role,
		CaptchaToken:  req.CaptchaToken,
		ReferralCode:  req.ReferralCode,
		SocialLinks:   req.SocialLinks,
		FollowerCount: req.FollowerCount,
	}
	
	// Register user
	user, err := h.authService.Register(c.Request.Context(), serviceReq, tenantID)
	if err != nil {
		h.logger.Error("Failed to register user", err)
		
		// Check for specific errors
		if strings.Contains(err.Error(), "already exists") {
			response.Error(c, http.StatusConflict, err.Error(), nil)
			return
		}
		
		response.Error(c, http.StatusInternalServerError, "Failed to register user", err)
		return
	}
	
	// Prepare response
	resp := response.UserResponse{
		ID:           user.ID.String(),
		Email:        user.Email,
		FullName:     user.FullName,
		Role:         string(user.Role),
		Status:       user.Status,
		ReferralCode: user.ReferralCode,
		CreatedAt:    user.CreatedAt,
	}
	
	// Return success response
	response.Success(c, http.StatusCreated, "User registered successfully", resp)
}

// Login handles user login
func (h *AuthHandler) Login(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Failed to bind login request", err)
		response.Error(c, http.StatusBadRequest, "Invalid request", err)
		return
	}
	
	// Convert request to service request
	serviceReq := service.AuthRequest{
		Email:        req.Email,
		Password:     req.Password,
		CaptchaToken: req.CaptchaToken,
	}
	
	// Login user
	tokenResp, user, err := h.authService.Login(c.Request.Context(), serviceReq)
	if err != nil {
		h.logger.Error("Failed to login user", err)
		
		// Check for specific errors
		if strings.Contains(err.Error(), "invalid email or password") {
			response.Error(c, http.StatusUnauthorized, "Invalid email or password", nil)
			return
		}
		
		if strings.Contains(err.Error(), "pending approval") {
			response.Error(c, http.StatusForbidden, "Your account is pending approval", nil)
			return
		}
		
		if strings.Contains(err.Error(), "inactive") {
			response.Error(c, http.StatusForbidden, "Your account is inactive", nil)
			return
		}
		
		response.Error(c, http.StatusInternalServerError, "Failed to login", err)
		return
	}
	
	// Prepare response
	resp := response.LoginResponse{
		Token: response.TokenResponse{
			AccessToken:  tokenResp.AccessToken,
			RefreshToken: tokenResp.RefreshToken,
			TokenType:    tokenResp.TokenType,
			ExpiresIn:    tokenResp.ExpiresIn,
		},
		User: response.UserResponse{
			ID:           user.ID.String(),
			Email:        user.Email,
			FullName:     user.FullName,
			Role:         string(user.Role),
			Status:       user.Status,
			ReferralCode: user.ReferralCode,
			CreatedAt:    user.CreatedAt,
		},
	}
	
	// Return success response
	response.Success(c, http.StatusOK, "Login successful", resp)
}

// RefreshToken handles token refresh
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req request.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Failed to bind refresh token request", err)
		response.Error(c, http.StatusBadRequest, "Invalid request", err)
		return
	}
	
	// Refresh token
	tokenResp, err := h.authService.RefreshToken(c.Request.Context(), req.RefreshToken)
	if err != nil {
		h.logger.Error("Failed to refresh token", err)
		response.Error(c, http.StatusUnauthorized, "Invalid refresh token", nil)
		return
	}
	
	// Prepare response
	resp := response.TokenResponse{
		AccessToken:  tokenResp.AccessToken,
		RefreshToken: tokenResp.RefreshToken,
		TokenType:    tokenResp.TokenType,
		ExpiresIn:    tokenResp.ExpiresIn,
	}
	
	// Return success response
	response.Success(c, http.StatusOK, "Token refreshed successfully", resp)
}

// Logout handles user logout
func (h *AuthHandler) Logout(c *gin.Context) {
	// Get token ID from context (set by middleware)
	tokenID := c.GetString("tokenID")
	
	// Get user ID from context
	userIDStr := c.GetString("userID")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		h.logger.Error("Invalid user ID", err)
		response.Error(c, http.StatusBadRequest, "Invalid user ID", err)
		return
	}
	
	// Logout user
	err = h.authService.Logout(c.Request.Context(), userID, tokenID)
	if err != nil {
		h.logger.Error("Failed to logout user", err)
		response.Error(c, http.StatusInternalServerError, "Failed to logout", err)
		return
	}
	
	// Return success response
	response.Success(c, http.StatusOK, "Logout successful", nil)
}

// GoogleLogin handles Google OAuth login
func (h *AuthHandler) GoogleLogin(c *gin.Context) {
	var req request.GoogleLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Failed to bind Google login request", err)
		response.Error(c, http.StatusBadRequest, "Invalid request", err)
		return
	}
	
	// Get tenant ID from context (set by middleware)
	tenantIDStr := c.GetString("tenantID")
	if tenantIDStr == "" {
		// If not set, use a default tenant ID
		tenantIDStr = "00000000-0000-0000-0000-000000000000"
	}
	
	tenantID, err := uuid.Parse(tenantIDStr)
	if err != nil {
		h.logger.Error("Invalid tenant ID", err)
		response.Error(c, http.StatusBadRequest, "Invalid tenant ID", err)
		return
	}
	
	// Convert request to service request
	serviceReq := service.GoogleOAuthRequest{
		Code:        req.Code,
		RedirectURI: req.RedirectURI,
	}
	
	// Login with Google
	tokenResp, user, err := h.authService.GoogleOAuthLogin(c.Request.Context(), serviceReq, tenantID)
	if err != nil {
		h.logger.Error("Failed to login with Google", err)
		response.Error(c, http.StatusInternalServerError, "Failed to login with Google", err)
		return
	}
	
	// Prepare response
	resp := response.LoginResponse{
		Token: response.TokenResponse{
			AccessToken:  tokenResp.AccessToken,
			RefreshToken: tokenResp.RefreshToken,
			TokenType:    tokenResp.TokenType,
			ExpiresIn:    tokenResp.ExpiresIn,
		},
		User: response.UserResponse{
			ID:           user.ID.String(),
			Email:        user.Email,
			FullName:     user.FullName,
			Role:         string(user.Role),
			Status:       user.Status,
			ReferralCode: user.ReferralCode,
			CreatedAt:    user.CreatedAt,
		},
	}
	
	// Return success response
	response.Success(c, http.StatusOK, "Google login successful", resp)
}