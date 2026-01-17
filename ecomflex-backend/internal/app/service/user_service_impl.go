package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/naresh6454/ecomflex-backend/internal/domain/entity"
	"github.com/naresh6454/ecomflex-backend/internal/domain/repository"
	"github.com/naresh6454/ecomflex-backend/internal/domain/service"
	"github.com/naresh6454/ecomflex-backend/internal/util"
	loggerPkg "github.com/naresh6454/ecomflex-backend/internal/util/logger"
)

// UserServiceImpl implements UserService interface
type UserServiceImpl struct {
	userRepo repository.UserRepository
	logger   loggerPkg.Logger
}

// NewUserService creates a new UserServiceImpl
func NewUserService(userRepo repository.UserRepository, logger loggerPkg.Logger) service.UserService {
	return &UserServiceImpl{
		userRepo: userRepo,
		logger:   logger,
	}
}

// GetUserByID gets a user by ID
func (s *UserServiceImpl) GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	return s.userRepo.GetUserByID(ctx, id)
}

// UpdateUser updates a user's information
func (s *UserServiceImpl) UpdateUser(ctx context.Context, id uuid.UUID, req service.UserUpdateRequest) (*entity.User, error) {
	// Get existing user
	user, err := s.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// Update user fields
	if req.FullName != "" {
		user.FullName = req.FullName
	}
	if req.Phone != "" {
		user.Phone = req.Phone
	}
	if req.ProfilePicture != "" {
		user.ProfilePicture = req.ProfilePicture
	}
	if req.SocialLinks != nil {
		user.SocialLinks = req.SocialLinks
	}

	// Set updated_at
	user.UpdatedAt = time.Now()

	// Save user to database
	err = s.userRepo.UpdateUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return user, nil
}

// UpdateInfluencerStatus updates an influencer's approval status
func (s *UserServiceImpl) UpdateInfluencerStatus(ctx context.Context, id uuid.UUID, req service.InfluencerStatusUpdateRequest) error {
	// Get existing user
	user, err := s.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	// Verify the user is an influencer
	if user.Role != entity.RoleInfluencer {
		return errors.New("user is not an influencer")
	}

	// Update status based on request
	switch req.Status {
	case "active":
		user.SetActive()
	case "rejected":
		user.SetRejected()
		// TODO: Send rejection email with reason
	default:
		return errors.New("invalid status")
	}

	// Save user to database
	err = s.userRepo.UpdateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to update influencer status: %w", err)
	}

	return nil
}

// GetInfluencersByStatus gets influencers by status with filters
func (s *UserServiceImpl) GetInfluencersByStatus(ctx context.Context, status, search, fromDate string, limit, offset int) ([]*entity.User, int, error) {
	// If status is empty or "all", set it to empty to get all influencers
	if status == "all" {
		status = ""
	}

	// Get influencers from repository
	return s.userRepo.GetInfluencersByStatus(ctx, status, limit, offset)
}

// GetPendingInfluencers gets all pending influencers with pagination
func (s *UserServiceImpl) GetPendingInfluencers(ctx context.Context, limit, offset int) ([]*entity.User, int, error) {
	// Use "pending_approval" as the status to get pending influencers
	return s.userRepo.GetInfluencersByStatus(ctx, "pending_approval", limit, offset)
}

// IsReferralCodeValid checks if a referral code is valid
func (s *UserServiceImpl) IsReferralCodeValid(ctx context.Context, referralCode string) (bool, error) {
	// Get user by referral code
	user, err := s.userRepo.GetUserByReferralCode(ctx, referralCode)
	if err != nil {
		if strings.Contains(err.Error(), "user not found") {
			return false, nil
		}
		return false, fmt.Errorf("failed to check referral code: %w", err)
	}

	// Check if user is active
	return user.IsActive(), nil
}

// GenerateUniqueReferralCode generates a unique referral code
func (s *UserServiceImpl) GenerateUniqueReferralCode(ctx context.Context, prefix string) (string, error) {
	// Generate a base code using the prefix
	baseCode := util.GenerateReferralCode(prefix)

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