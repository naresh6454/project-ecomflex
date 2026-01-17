package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/naresh6454/ecomflex-backend/config"
	"github.com/naresh6454/ecomflex-backend/internal/domain/entity"
	"github.com/naresh6454/ecomflex-backend/internal/domain/repository"
	"github.com/naresh6454/ecomflex-backend/internal/domain/service"
	"github.com/naresh6454/ecomflex-backend/internal/infrastructure/auth"
	"github.com/naresh6454/ecomflex-backend/internal/util"
)

// AuthServiceImpl implements AuthService interface
type AuthServiceImpl struct {
	userRepo         repository.UserRepository
	authRepo         repository.AuthRepository
	tenantRepo       repository.TenantRepository // Added tenant repository
	jwtProvider      *auth.JWTProvider
	googleOAuthCfg   config.OAuthConfig
	captchaCfg       config.CaptchaConfig
	httpClient       *http.Client
}

// NewAuthService creates a new AuthServiceImpl
func NewAuthService(
	userRepo repository.UserRepository,
	authRepo repository.AuthRepository,
	tenantRepo repository.TenantRepository, // Added tenant repository
	jwtProvider *auth.JWTProvider,
	googleOAuthCfg config.OAuthConfig,
	captchaCfg config.CaptchaConfig,
) service.AuthService {
	return &AuthServiceImpl{
		userRepo:       userRepo,
		authRepo:       authRepo,
		tenantRepo:     tenantRepo, // Set tenant repository
		jwtProvider:    jwtProvider,
		googleOAuthCfg: googleOAuthCfg,
		captchaCfg:     captchaCfg,
		httpClient:     &http.Client{Timeout: 10 * time.Second},
	}
}

// Register registers a new user
func (s *AuthServiceImpl) Register(ctx context.Context, req service.RegisterRequest, tenantID uuid.UUID) (*entity.User, error) {
	// Skip captcha validation for development
	// Removed validation code here
	
	// Check if email is already taken
	emailTaken, err := s.userRepo.IsEmailTaken(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to check email: %w", err)
	}
	
	if emailTaken {
		return nil, errors.New("email already exists")
	}
	
	// Hash password
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	
	var user *entity.User
	
	// Create user based on role
	switch req.Role {
	case "influencer":
		// For influencers, CREATE A NEW TENANT instead of using the default one
		// Generate a tenant name based on the influencer's name
		tenantName := fmt.Sprintf("%s's Store", req.FullName)
		
		// Generate a unique domain from the influencer's name
		domainBase := strings.ToLower(strings.ReplaceAll(req.FullName, " ", "-"))
		domain := fmt.Sprintf("%s.ecomflex.com", domainBase)
		
		// Create a new tenant entity
		newTenant := entity.NewTenant(tenantName, domain, "free")
		
		// Save tenant to database
		if err := s.tenantRepo.CreateTenant(ctx, newTenant); err != nil {
			return nil, fmt.Errorf("failed to create tenant: %w", err)
		}
		
		// Use the new tenant ID instead of the default one
		tenantID = newTenant.ID
		
		// For influencers, validate referral code
		if req.ReferralCode != "" {
			referralCodeTaken, err := s.userRepo.IsReferralCodeTaken(ctx, req.ReferralCode)
			if err != nil {
				return nil, fmt.Errorf("failed to check referral code: %w", err)
			}
			
			if referralCodeTaken {
				return nil, errors.New("referral code already exists")
			}
		} else {
			// Generate a unique referral code if not provided
			req.ReferralCode, err = s.generateUniqueReferralCode(ctx, req.FullName)
			if err != nil {
				return nil, fmt.Errorf("failed to generate referral code: %w", err)
			}
		}
		
		// Create influencer user with the NEW tenant ID
		user = entity.NewInfluencer(
			tenantID,
			req.Email,
			hashedPassword,
			req.FullName,
			req.Phone,
			req.ReferralCode,
			req.SocialLinks,
			req.FollowerCount,
		)
	case "super_admin":
		// Create super admin user
		user = entity.NewUser(
			tenantID,
			req.Email,
			hashedPassword,
			req.FullName,
			entity.RoleSuperAdmin,
			req.Phone,
		)
		// Super admins are always active
		user.SetActive()
	default:
		// Create public user
		user = entity.NewUser(
			tenantID,
			req.Email,
			hashedPassword,
			req.FullName,
			entity.RolePublic,
			req.Phone,
		)
	}
	
	// Save user to database
	err = s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	
	return user, nil
}

// Login authenticates a user and returns tokens
func (s *AuthServiceImpl) Login(ctx context.Context, req service.AuthRequest) (*service.TokenResponse, *entity.User, error) {
	// Hard-coded superadmin check
	if req.Email == "admin@ecomflex.com" && req.Password == "Admin@123" {
		// Create a temporary superadmin user for token generation
		// This user is not saved to the database
		superadminUser := &entity.User{
			ID:           uuid.New(),
			TenantID:     uuid.MustParse("00000000-0000-0000-0000-000000000000"),
			Email:        "admin@ecomflex.com",
			PasswordHash: "", // Not needed as we've already verified the password
			FullName:     "System Administrator",
			Role:         entity.RoleSuperAdmin,
			Phone:        "",
			Status:       "active",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		// Generate tokens for the superadmin
		tokens, err := s.generateTokens(ctx, superadminUser)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to generate tokens for superadmin: %w", err)
		}
		
		return tokens, superadminUser, nil
	}

	// Regular login flow
	// Get user by email
	user, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, nil, errors.New("invalid email or password")
	}
	
	// Check password
	if !util.CheckPasswordHash(req.Password, user.PasswordHash) {
		return nil, nil, errors.New("invalid email or password")
	}
	
	// Check if user is active
	if !user.IsActive() {
		// For influencers who are pending approval
		if user.Role == entity.RoleInfluencer && user.IsPendingApproval() {
			return nil, nil, errors.New("your account is pending approval")
		}
		
		return nil, nil, errors.New("your account is inactive")
	}
	
	// Generate tokens
	tokens, err := s.generateTokens(ctx, user)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate tokens: %w", err)
	}
	
	return tokens, user, nil
}

// RefreshToken refreshes an access token using a refresh token
func (s *AuthServiceImpl) RefreshToken(ctx context.Context, refreshToken string) (*service.TokenResponse, error) {
	// Extract token ID from refresh token
	tokenID := refreshToken[:36] // UUID is 36 chars
	
	// Get the stored refresh token
	storedToken, err := s.authRepo.GetRefreshToken(ctx, tokenID)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}
	
	// Verify the token matches
	if storedToken != refreshToken {
		return nil, errors.New("invalid refresh token")
	}
	
	// Delete the used refresh token
	err = s.authRepo.DeleteRefreshToken(ctx, tokenID)
	if err != nil {
		return nil, fmt.Errorf("failed to delete refresh token: %w", err)
	}
	
	// Parse claims from token ID to get user ID
	tokenParts := refreshToken[37:] // Skip UUID and separator
	userID, err := uuid.Parse(tokenParts[:36]) // UUID is 36 chars
	if err != nil {
		return nil, errors.New("invalid refresh token format")
	}
	
	// Get user
	user, err := s.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, errors.New("user not found")
	}
	
	// Check if user is active
	if !user.IsActive() {
		return nil, errors.New("your account is inactive")
	}
	
	// Generate new tokens
	tokens, err := s.generateTokens(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to generate tokens: %w", err)
	}
	
	return tokens, nil
}

// Logout logs out a user and revokes their tokens
func (s *AuthServiceImpl) Logout(ctx context.Context, userID uuid.UUID, tokenID string) error {
	// Get token expiration time (we need to store it in revocation list until it expires)
	// For now, we'll use a default expiration of 24 hours
	expiresAt := time.Now().Add(24 * time.Hour)
	
	// Revoke the access token
	err := s.authRepo.RevokeToken(ctx, tokenID, expiresAt)
	if err != nil {
		return fmt.Errorf("failed to revoke token: %w", err)
	}
	
	// Delete all refresh tokens for this user
	err = s.authRepo.DeleteAllUserRefreshTokens(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to delete refresh tokens: %w", err)
	}
	
	return nil
}

// GoogleOAuthLogin handles Google OAuth login
func (s *AuthServiceImpl) GoogleOAuthLogin(ctx context.Context, req service.GoogleOAuthRequest, tenantID uuid.UUID) (*service.TokenResponse, *entity.User, error) {
	// TODO: Implement Google OAuth login
	// This will involve:
	// 1. Exchanging the code for tokens
	// 2. Getting user info from Google
	// 3. Finding or creating the user in our system
	// 4. Generating tokens
	
	return nil, nil, errors.New("google oauth login not implemented")
}

// ValidateCaptcha validates a reCAPTCHA token
func (s *AuthServiceImpl) ValidateCaptcha(ctx context.Context, captchaToken string, remoteIP string) (bool, error) {
	// Always return true for now to bypass captcha validation
	return true, nil
}

// GetUserFromToken retrieves a user from a JWT token
func (s *AuthServiceImpl) GetUserFromToken(ctx context.Context, token string) (*entity.User, error) {
	// Special case for superadmin token
	claims, err := s.jwtProvider.ValidateToken(token)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}
	
	// Check if token is for superadmin
	if claims.Email == "admin@ecomflex.com" && claims.Role == string(entity.RoleSuperAdmin) {
		// Return a superadmin user without checking the database
		return &entity.User{
			ID:           uuid.MustParse(claims.UserID),
			TenantID:     uuid.MustParse(claims.TenantID),
			Email:        claims.Email,
			FullName:     "System Administrator",
			Role:         entity.RoleSuperAdmin,
			Status:       "active",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}, nil
	}
	
	// Check if token is revoked
	revoked, err := s.authRepo.IsTokenRevoked(ctx, claims.TokenID)
	if err != nil {
		return nil, fmt.Errorf("failed to check if token is revoked: %w", err)
	}
	
	if revoked {
		return nil, errors.New("token has been revoked")
	}
	
	// Parse user ID
	userID, err := uuid.Parse(claims.UserID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID in token: %w", err)
	}
	
	// Get user
	user, err := s.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, errors.New("user not found")
	}
	
	return user, nil
}

// generateTokens generates access and refresh tokens for a user
func (s *AuthServiceImpl) generateTokens(ctx context.Context, user *entity.User) (*service.TokenResponse, error) {
	// Generate access token
	accessToken, tokenID, expiresAt, err := s.jwtProvider.GenerateAccessToken(user)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}
	
	// Generate refresh token
	refreshTokenValue, err := s.jwtProvider.GenerateRefreshToken()
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}
	
	// Create a composite refresh token: tokenID + _ + userID + _ + refreshTokenValue
	// This way we can extract the tokenID and userID from the token itself
	refreshToken := fmt.Sprintf("%s_%s_%s", tokenID, user.ID.String(), refreshTokenValue)
	
	// Calculate refresh token expiration
	refreshTokenExp := time.Now().Add(24 * 7 * time.Hour) // 7 days
	
	// Store refresh token
	err = s.authRepo.SaveRefreshToken(ctx, user.ID, tokenID, refreshToken, refreshTokenExp)
	if err != nil {
		return nil, fmt.Errorf("failed to save refresh token: %w", err)
	}
	
	// Create token response
	response := &service.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    int64(time.Until(expiresAt).Seconds()),
	}
	
	return response, nil
}

// generateUniqueReferralCode generates a unique referral code for a user
func (s *AuthServiceImpl) generateUniqueReferralCode(ctx context.Context, fullName string) (string, error) {
	// Generate a base code using the first few characters of the name
	baseCode := util.GenerateReferralCode(fullName)
	
	// Check if the code is already taken
	taken, err := s.userRepo.IsReferralCodeTaken(ctx, baseCode)
	if err != nil {
		return "", fmt.Errorf("failed to check referral code: %w", err)
	}
	
	// If not taken, return the base code
	if !taken {
		return baseCode, nil
	}
	
	// Try adding numbers until we find a unique code
	for i := 1; i < 1000; i++ {
		newCode := fmt.Sprintf("%s%d", baseCode, i)
		
		taken, err := s.userRepo.IsReferralCodeTaken(ctx, newCode)
		if err != nil {
			return "", fmt.Errorf("failed to check referral code: %w", err)
		}
		
		if !taken {
			return newCode, nil
		}
	}
	
	// If we can't find a unique code after 1000 tries, generate a random one
	return util.GenerateRandomReferralCode(), nil
}