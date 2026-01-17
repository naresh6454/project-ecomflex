package request

// RegisterRequest represents a user registration request
type RegisterRequest struct {
	Email         string   `json:"email" binding:"required,email"`
	Password      string   `json:"password" binding:"required,min=8"`
	FullName      string   `json:"full_name" binding:"required"`
	Phone         string   `json:"phone" binding:"omitempty"`
	Role          string   `json:"role" binding:"required,oneof=public influencer"`
	CaptchaToken  string   `json:"captcha_token"` // Removed required binding
	ReferralCode  string   `json:"referral_code" binding:"omitempty"`
	SocialLinks   []string `json:"social_links" binding:"omitempty"`
	FollowerCount int      `json:"follower_count" binding:"omitempty"`
}

// LoginRequest represents a user login request
type LoginRequest struct {
	Email        string `json:"email" binding:"required,email"`
	Password     string `json:"password" binding:"required"`
	CaptchaToken string `json:"captcha_token"` // Removed required binding
}

// RefreshTokenRequest represents a refresh token request
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// GoogleLoginRequest represents a Google OAuth login request
type GoogleLoginRequest struct {
	Code        string `json:"code" binding:"required"`
	RedirectURI string `json:"redirect_uri" binding:"required"`
}