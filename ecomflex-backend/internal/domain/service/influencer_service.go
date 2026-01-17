package service

import (
	"context"
)

// ReferralStats represents statistics for referrals
type ReferralStats struct {
	TotalReferrals   int     `json:"total_referrals"`
	ApprovedBookings int     `json:"approved_bookings"`
	PendingBookings  int     `json:"pending_bookings"`
	TotalEarnings    float64 `json:"total_earnings"`
}

// Referral represents a referral made by an influencer
type Referral struct {
	ID       string  `json:"id"`
	User     string  `json:"user"`
	Product  string  `json:"product"`
	Date     string  `json:"date"`
	Status   string  `json:"status"`
	Earnings float64 `json:"earnings"`
}

// DashboardData represents the data for an influencer's dashboard
type DashboardData struct {
	User            map[string]interface{} `json:"user"`
	ReferralCode    string                 `json:"referral_code"`
	ReferralLink    string                 `json:"referral_link"`
	Stats           ReferralStats          `json:"stats"`
	RecentReferrals []Referral             `json:"recent_referrals"`
}

// InfluencerService defines operations for managing influencer data
type InfluencerService interface {
	// GetReferralCode returns the influencer's referral code
	GetReferralCode(ctx context.Context, userID string, tenantID string) (string, error)
	
	// GetReferralStats returns the influencer's referral statistics
	GetReferralStats(ctx context.Context, userID string, tenantID string) (ReferralStats, error)
	
	// GetRecentReferrals returns the influencer's recent referrals
	GetRecentReferrals(ctx context.Context, userID string, tenantID string, limit int) ([]Referral, error)
	
	// GetDashboardData returns all data needed for the influencer dashboard
	GetDashboardData(ctx context.Context, userID string, tenantID string) (DashboardData, error)
}