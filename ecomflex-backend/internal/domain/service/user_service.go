package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/naresh6454/ecomflex-backend/internal/domain/entity"
)

// UserUpdateRequest represents a request to update a user
type UserUpdateRequest struct {
	FullName       string   `json:"full_name" binding:"omitempty"`
	Phone          string   `json:"phone" binding:"omitempty"`
	ProfilePicture string   `json:"profile_picture" binding:"omitempty"`
	SocialLinks    []string `json:"social_links" binding:"omitempty"`
}

// InfluencerStatusUpdateRequest represents a request to update influencer status
type InfluencerStatusUpdateRequest struct {
	Status string `json:"status" binding:"required,oneof=active rejected"`
	Reason string `json:"reason" binding:"omitempty"`
}

// UserService defines the interface for user service
type UserService interface {
	// GetUserByID gets a user by ID
	GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	
	// UpdateUser updates a user's information
	UpdateUser(ctx context.Context, id uuid.UUID, req UserUpdateRequest) (*entity.User, error)
	
	// UpdateInfluencerStatus updates an influencer's approval status
	UpdateInfluencerStatus(ctx context.Context, id uuid.UUID, req InfluencerStatusUpdateRequest) error
	
	// GetPendingInfluencers gets all pending influencers
	GetPendingInfluencers(ctx context.Context, limit, offset int) ([]*entity.User, int, error)
	
	// IsReferralCodeValid checks if a referral code is valid
	IsReferralCodeValid(ctx context.Context, referralCode string) (bool, error)
	
	// GenerateUniqueReferralCode generates a unique referral code
	GenerateUniqueReferralCode(ctx context.Context, prefix string) (string, error)
}