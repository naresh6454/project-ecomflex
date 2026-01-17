package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/naresh6454/ecomflex-backend/internal/domain/entity"
)

// UserRepository defines the interface for user repository operations
type UserRepository interface {
	// CreateUser creates a new user
	CreateUser(ctx context.Context, user *entity.User) error
	
	// GetUserByID gets a user by ID
	GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	
	// GetUserByEmail gets a user by email
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	
	// GetUserByReferralCode gets a user by referral code
	GetUserByReferralCode(ctx context.Context, referralCode string) (*entity.User, error)
	
	// UpdateUser updates a user
	UpdateUser(ctx context.Context, user *entity.User) error
	
	// DeleteUser deletes a user
	DeleteUser(ctx context.Context, id uuid.UUID) error
	
	// GetInfluencersByStatus gets influencers by status
	GetInfluencersByStatus(ctx context.Context, status string, limit, offset int) ([]*entity.User, int, error)
	
	// IsEmailTaken checks if an email is already taken
	IsEmailTaken(ctx context.Context, email string) (bool, error)
	
	// IsReferralCodeTaken checks if a referral code is already taken
	IsReferralCodeTaken(ctx context.Context, referralCode string) (bool, error)
}