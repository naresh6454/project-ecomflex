package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/naresh6454/ecomflex-backend/internal/domain/entity"
)

// AuthRequest represents a user authentication request
type AuthRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8"`
	CaptchaToken string `json:"captcha_token" binding:"required"`
}

// RegisterRequest represents a user registration request
type RegisterRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8"`
	FullName  string `json:"full_name" binding:"required"`
	Phone     string `json:"phone" binding:"omitempty"`
	Role      string `json:"role" binding:"required,oneof=public influencer"`
	CaptchaToken string `json:"captcha_token" binding:"required"`
	
	// Fields for influencer
	ReferralCode  string   `json:"referral_code" binding:"omitempty"`
	SocialLinks   []string `json:"social_links" binding:"omitempty"`
	FollowerCount int      `json:"follower_count" binding:"omitempty"`
}

// GoogleOAuthRequest represents a Google OAuth login request
type GoogleOAuthRequest struct {
	Code        string `json:"code" binding:"required"`
	RedirectURI string `json:"redirect_uri" binding:"required"`
}

// TokenResponse represents a token response
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"` // seconds
}

// AuthService defines the interface for authentication service
type AuthService interface {
	// Register registers a new user
	Register(ctx context.Context, req RegisterRequest, tenantID uuid.UUID) (*entity.User, error)
	
	// Login authenticates a user and returns tokens
	Login(ctx context.Context, req AuthRequest) (*TokenResponse, *entity.User, error)
	
	// RefreshToken refreshes an access token using a refresh token
	RefreshToken(ctx context.Context, refreshToken string) (*TokenResponse, error)
	
	// Logout logs out a user and revokes their tokens
	Logout(ctx context.Context, userID uuid.UUID, tokenID string) error
	
	// GoogleOAuthLogin handles Google OAuth login
	GoogleOAuthLogin(ctx context.Context, req GoogleOAuthRequest, tenantID uuid.UUID) (*TokenResponse, *entity.User, error)
	
	// ValidateCaptcha validates a reCAPTCHA token
	ValidateCaptcha(ctx context.Context, captchaToken string, remoteIP string) (bool, error)
	
	// GetUserFromToken retrieves a user from a JWT token
	GetUserFromToken(ctx context.Context, token string) (*entity.User, error)
}