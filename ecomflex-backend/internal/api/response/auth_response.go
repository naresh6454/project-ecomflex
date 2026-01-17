package response

import (
	"time"
)

// TokenResponse represents a token response
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"` // seconds
}

// UserResponse represents a user response
type UserResponse struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	FullName     string    `json:"full_name"`
	Role         string    `json:"role"`
	Status       string    `json:"status"`
	ReferralCode string    `json:"referral_code,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
}

// LoginResponse represents a login response
type LoginResponse struct {
	Token TokenResponse `json:"token"`
	User  UserResponse  `json:"user"`
}