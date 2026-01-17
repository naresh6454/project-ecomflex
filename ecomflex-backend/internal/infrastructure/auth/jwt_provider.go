package auth

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/naresh6454/ecomflex-backend/config"
	"github.com/naresh6454/ecomflex-backend/internal/domain/entity"
)

// CustomClaims represents custom JWT claims
type CustomClaims struct {
	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	TenantID string `json:"tenant_id"`
	TokenID  string `json:"token_id"`
	jwt.RegisteredClaims
}

// JWTProvider provides JWT token generation and validation
type JWTProvider struct {
	config *config.JWTConfig
}

// NewJWTProvider creates a new JWT provider
func NewJWTProvider(config *config.JWTConfig) *JWTProvider {
	return &JWTProvider{
		config: config,
	}
}

// GenerateAccessToken generates a new JWT access token for a user
func (p *JWTProvider) GenerateAccessToken(user *entity.User) (string, string, time.Time, error) {
	// Generate a unique token ID
	tokenID := uuid.New().String()
	
	// Calculate expiration time
	expirationTime := time.Now().Add(p.config.AccessTokenExp)
	
	// Create claims
	claims := CustomClaims{
		UserID:   user.ID.String(),
		Email:    user.Email,
		Role:     string(user.Role),
		TenantID: user.TenantID.String(),
		TokenID:  tokenID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Subject:   user.ID.String(),
		},
	}
	
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	// Sign token
	tokenString, err := token.SignedString([]byte(p.config.Secret))
	if err != nil {
		return "", "", time.Time{}, err
	}
	
	return tokenString, tokenID, expirationTime, nil
}

// GenerateRefreshToken generates a new refresh token
func (p *JWTProvider) GenerateRefreshToken() (string, error) {
	b := make([]byte, p.config.RefreshTokenSize)
	_, err := p.randomReader().Read(b)
	if err != nil {
		return "", err
	}
	
	return fmt.Sprintf("%x", b), nil
}

// ValidateToken validates a JWT token
func (p *JWTProvider) ValidateToken(tokenString string) (*CustomClaims, error) {
	// Parse token
	token, err := jwt.ParseWithClaims(
		tokenString,
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			// Validate signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(p.config.Secret), nil
		},
	)
	
	if err != nil {
		return nil, err
	}
	
	// Validate claims
	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	
	return claims, nil
}

// randomReader is a helper function to get a reader for random bytes
// This is extracted to make testing easier (can be mocked)
func (p *JWTProvider) randomReader() io.Reader {
	return rand.Reader
}