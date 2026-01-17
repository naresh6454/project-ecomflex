package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// AuthRepository defines the interface for authentication repository operations
type AuthRepository interface {
	// SaveRefreshToken stores a refresh token in Redis
	SaveRefreshToken(ctx context.Context, userID uuid.UUID, tokenID string, token string, expiresAt time.Time) error
	
	// GetRefreshToken retrieves a refresh token from Redis
	GetRefreshToken(ctx context.Context, tokenID string) (string, error)
	
	// DeleteRefreshToken deletes a refresh token from Redis
	DeleteRefreshToken(ctx context.Context, tokenID string) error
	
	// DeleteAllUserRefreshTokens deletes all refresh tokens for a user
	DeleteAllUserRefreshTokens(ctx context.Context, userID uuid.UUID) error
	
	// IsTokenRevoked checks if a token has been revoked
	IsTokenRevoked(ctx context.Context, tokenID string) (bool, error)
	
	// RevokeToken adds a token to the revocation list
	RevokeToken(ctx context.Context, tokenID string, expiresAt time.Time) error
}