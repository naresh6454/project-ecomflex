package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/naresh6454/ecomflex-backend/internal/domain/repository"
)

const (
	// Keys for Redis
	refreshTokenKeyPrefix = "refresh_token:"
	revokedTokenKeyPrefix = "revoked_token:"
	userTokensKeyPrefix   = "user_tokens:"
)

// RedisAuthRepository implements AuthRepository interface using Redis
type RedisAuthRepository struct {
	redis *redis.Client
}

// NewRedisAuthRepository creates a new RedisAuthRepository
func NewRedisAuthRepository(redis *redis.Client) repository.AuthRepository {
	return &RedisAuthRepository{
		redis: redis,
	}
}

// SaveRefreshToken stores a refresh token in Redis
func (r *RedisAuthRepository) SaveRefreshToken(ctx context.Context, userID uuid.UUID, tokenID string, token string, expiresAt time.Time) error {
	// Create the refresh token key
	refreshTokenKey := refreshTokenKeyPrefix + tokenID
	
	// Set refresh token with expiration
	duration := time.Until(expiresAt)
	err := r.redis.Set(ctx, refreshTokenKey, token, duration).Err()
	if err != nil {
		return fmt.Errorf("failed to save refresh token: %w", err)
	}
	
	// Add token to user's token set
	userTokensKey := userTokensKeyPrefix + userID.String()
	err = r.redis.SAdd(ctx, userTokensKey, tokenID).Err()
	if err != nil {
		return fmt.Errorf("failed to add token to user tokens: %w", err)
	}
	
	// Set expiration on the user tokens set too (for cleanup)
	err = r.redis.ExpireAt(ctx, userTokensKey, expiresAt).Err()
	if err != nil {
		return fmt.Errorf("failed to set expiration on user tokens: %w", err)
	}
	
	return nil
}

// GetRefreshToken retrieves a refresh token from Redis
func (r *RedisAuthRepository) GetRefreshToken(ctx context.Context, tokenID string) (string, error) {
	refreshTokenKey := refreshTokenKeyPrefix + tokenID
	token, err := r.redis.Get(ctx, refreshTokenKey).Result()
	if err != nil {
		if err == redis.Nil {
			return "", fmt.Errorf("refresh token not found")
		}
		return "", fmt.Errorf("failed to get refresh token: %w", err)
	}
	
	return token, nil
}

// DeleteRefreshToken deletes a refresh token from Redis
func (r *RedisAuthRepository) DeleteRefreshToken(ctx context.Context, tokenID string) error {
	refreshTokenKey := refreshTokenKeyPrefix + tokenID
	err := r.redis.Del(ctx, refreshTokenKey).Err()
	if err != nil {
		return fmt.Errorf("failed to delete refresh token: %w", err)
	}
	
	// Note: We don't remove it from the user's token set here for efficiency,
	// as it would require knowing the userID. The tokens will be cleaned up
	// when they expire.
	
	return nil
}

// DeleteAllUserRefreshTokens deletes all refresh tokens for a user
func (r *RedisAuthRepository) DeleteAllUserRefreshTokens(ctx context.Context, userID uuid.UUID) error {
	userTokensKey := userTokensKeyPrefix + userID.String()
	
	// Get all token IDs for this user
	tokenIDs, err := r.redis.SMembers(ctx, userTokensKey).Result()
	if err != nil {
		return fmt.Errorf("failed to get user tokens: %w", err)
	}
	
	// If there are no tokens, we're done
	if len(tokenIDs) == 0 {
		return nil
	}
	
	// Create a pipeline to efficiently delete all tokens
	pipe := r.redis.Pipeline()
	
	// Delete each refresh token
	for _, tokenID := range tokenIDs {
		refreshTokenKey := refreshTokenKeyPrefix + tokenID
		pipe.Del(ctx, refreshTokenKey)
	}
	
	// Delete the user tokens set
	pipe.Del(ctx, userTokensKey)
	
	// Execute the pipeline
	_, err = pipe.Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete user tokens: %w", err)
	}
	
	return nil
}

// IsTokenRevoked checks if a token has been revoked
func (r *RedisAuthRepository) IsTokenRevoked(ctx context.Context, tokenID string) (bool, error) {
	revokedTokenKey := revokedTokenKeyPrefix + tokenID
	exists, err := r.redis.Exists(ctx, revokedTokenKey).Result()
	if err != nil {
		return false, fmt.Errorf("failed to check if token is revoked: %w", err)
	}
	
	return exists > 0, nil
}

// RevokeToken adds a token to the revocation list
func (r *RedisAuthRepository) RevokeToken(ctx context.Context, tokenID string, expiresAt time.Time) error {
	revokedTokenKey := revokedTokenKeyPrefix + tokenID
	duration := time.Until(expiresAt)
	
	// We just need to set a key with the token ID, the value doesn't matter
	err := r.redis.Set(ctx, revokedTokenKey, "1", duration).Err()
	if err != nil {
		return fmt.Errorf("failed to revoke token: %w", err)
	}
	
	return nil
}