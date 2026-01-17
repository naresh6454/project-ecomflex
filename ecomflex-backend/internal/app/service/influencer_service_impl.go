package service

import (
	"context"
	"fmt"
	"time"
	
	"github.com/google/uuid"
	"github.com/naresh6454/ecomflex-backend/internal/domain/repository"
	"github.com/naresh6454/ecomflex-backend/internal/domain/service"
	"github.com/naresh6454/ecomflex-backend/internal/util/logger"
)

// InfluencerServiceImpl implements the InfluencerService interface
type InfluencerServiceImpl struct {
	userRepo      repository.UserRepository
	referralRepo  repository.ReferralRepository
	logger        logger.Logger
	baseURL       string
}

// NewInfluencerService creates a new InfluencerService
func NewInfluencerService(
	userRepo repository.UserRepository, 
	referralRepo repository.ReferralRepository, 
	logger logger.Logger, 
	baseURL string,
) service.InfluencerService {
	return &InfluencerServiceImpl{
		userRepo:     userRepo,
		referralRepo: referralRepo,
		logger:       logger,
		baseURL:      baseURL,
	}
}

// GetReferralCode returns the influencer's referral code
func (s *InfluencerServiceImpl) GetReferralCode(ctx context.Context, userID string, tenantID string) (string, error) {
	// Convert string ID to UUID
	id, err := uuid.Parse(userID)
	if err != nil {
		errMsg := "Invalid user ID format"
		s.logger.Error(errMsg, err)
		return "", fmt.Errorf("%s: %w", errMsg, err)
	}
	
	// Get the user to ensure they exist and are an influencer
	user, err := s.userRepo.GetUserByID(ctx, id)
	if err != nil {
		errMsg := "Failed to find user"
		s.logger.Error(errMsg, err)
		return "", fmt.Errorf("%s: %w", errMsg, err)
	}
	
	// Ensure the user belongs to the correct tenant
	if user.TenantID.String() != tenantID {
		errMsg := "User not found in this tenant"
		s.logger.Error(errMsg, fmt.Errorf("tenant mismatch"))
		return "", fmt.Errorf("%s", errMsg)
	}
	
	if user.ReferralCode != "" {
		return user.ReferralCode, nil
	}
	
	// Generate a referral code using a unique identifier
	username := user.Email[:4] // Use first part of email as fallback
	referralCode := fmt.Sprintf("%s%d", username, time.Now().UnixNano()%10000)
	
	// Update the user with the new referral code
	user.ReferralCode = referralCode
	err = s.userRepo.UpdateUser(ctx, user)
	if err != nil {
		errMsg := "Failed to update referral code"
		s.logger.Error(errMsg, err)
		return "", fmt.Errorf("%s: %w", errMsg, err)
	}
	
	return referralCode, nil
}

// GetReferralStats returns the influencer's referral statistics
func (s *InfluencerServiceImpl) GetReferralStats(ctx context.Context, userID string, tenantID string) (service.ReferralStats, error) {
	// Pass tenantID to repository for proper filtering
	stats, err := s.referralRepo.GetReferralStats(ctx, userID, tenantID)
	if err != nil {
		errMsg := "Failed to get referral stats"
		s.logger.Error(errMsg, err)
		return service.ReferralStats{}, fmt.Errorf("%s: %w", errMsg, err)
	}
	
	return service.ReferralStats{
		TotalReferrals:   stats.TotalReferrals,
		ApprovedBookings: stats.ApprovedBookings,
		PendingBookings:  stats.PendingBookings,
		TotalEarnings:    stats.TotalEarnings,
	}, nil
}

// GetRecentReferrals returns the influencer's recent referrals
func (s *InfluencerServiceImpl) GetRecentReferrals(ctx context.Context, userID string, tenantID string, limit int) ([]service.Referral, error) {
	// Pass tenantID to repository for proper filtering
	records, err := s.referralRepo.GetRecentReferrals(ctx, userID, tenantID, limit)
	if err != nil {
		errMsg := "Failed to get recent referrals"
		s.logger.Error(errMsg, err)
		return nil, fmt.Errorf("%s: %w", errMsg, err)
	}
	
	referrals := make([]service.Referral, len(records))
	for i, record := range records {
		referrals[i] = service.Referral{
			ID:       record.ID,
			User:     record.User,
			Product:  record.Product,
			Date:     record.Date,
			Status:   record.Status,
			Earnings: record.Earnings,
		}
	}
	
	return referrals, nil
}

// GetDashboardData returns all data needed for the influencer dashboard
func (s *InfluencerServiceImpl) GetDashboardData(ctx context.Context, userID string, tenantID string) (service.DashboardData, error) {
	// Convert string ID to UUID
	id, err := uuid.Parse(userID)
	if err != nil {
		errMsg := "Invalid user ID format"
		s.logger.Error(errMsg, err)
		return service.DashboardData{}, fmt.Errorf("%s: %w", errMsg, err)
	}
	
	// Get user data from the database
	user, err := s.userRepo.GetUserByID(ctx, id)
	if err != nil {
		errMsg := "Failed to find user"
		s.logger.Error(errMsg, err)
		return service.DashboardData{}, fmt.Errorf("%s: %w", errMsg, err)
	}
	
	// Ensure the user belongs to the correct tenant
	if user.TenantID.String() != tenantID {
		errMsg := "User not found in this tenant"
		s.logger.Error(errMsg, fmt.Errorf("tenant mismatch"))
		return service.DashboardData{}, fmt.Errorf("%s", errMsg)
	}
	
	// Get referral code - use existing one or generate new
	referralCode := user.ReferralCode
	if referralCode == "" {
		var codeErr error
		referralCode, codeErr = s.GetReferralCode(ctx, userID, tenantID)
		if codeErr != nil {
			errMsg := "Failed to get referral code"
			s.logger.Error(errMsg, codeErr)
			return service.DashboardData{}, fmt.Errorf("%s: %w", errMsg, codeErr)
		}
	}
	
	// Get referral stats
	stats, err := s.GetReferralStats(ctx, userID, tenantID)
	if err != nil {
		// If stats fail, create empty stats rather than failing completely
		s.logger.Warn("Failed to get referral stats, using empty stats", err)
		stats = service.ReferralStats{
			TotalReferrals:   0,
			ApprovedBookings: 0,
			PendingBookings:  0,
			TotalEarnings:    0,
		}
	}
	
	// Get recent referrals
	recentReferrals, err := s.GetRecentReferrals(ctx, userID, tenantID, 5)
	if err != nil {
		// If recent referrals fail, use empty list rather than failing completely
		s.logger.Warn("Failed to get recent referrals, using empty list", err)
		recentReferrals = []service.Referral{}
	}
	
	// Format user data for response - using fields that actually exist in your User entity
	userData := map[string]interface{}{
		"id":            user.ID.String(),
		"email":         user.Email,
		"fullName":      user.FullName,
		"phone":         user.Phone,
		"role":          user.Role,
		"createdAt":     user.CreatedAt.Format(time.RFC3339),
		"lastLogin":     user.UpdatedAt.Format(time.RFC3339),
		"referralCode":  user.ReferralCode,
		"tenantID":      user.TenantID.String(), // Add tenant ID for debugging
	}
	
	// Generate referral link
	referralLink := fmt.Sprintf("%s/ref/%s", s.baseURL, referralCode)
	
	return service.DashboardData{
		User:            userData,
		ReferralCode:    referralCode,
		ReferralLink:    referralLink,
		Stats:           stats,
		RecentReferrals: recentReferrals,
	}, nil
}