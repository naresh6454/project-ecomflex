package repository

import (
	"context"
)

// ReferralStats represents statistics for referrals
type ReferralStats struct {
	TotalReferrals   int     `json:"total_referrals" db:"total_referrals"`
	ApprovedBookings int     `json:"approved_bookings" db:"approved_bookings"`
	PendingBookings  int     `json:"pending_bookings" db:"pending_bookings"`
	TotalEarnings    float64 `json:"total_earnings" db:"total_earnings"`
}

// ReferralRecord represents a record of a referral
type ReferralRecord struct {
	ID       string  `json:"id" db:"id"`
	User     string  `json:"user" db:"user"`
	Product  string  `json:"product" db:"product"`
	Date     string  `json:"date" db:"date"`
	Status   string  `json:"status" db:"status"`
	Earnings float64 `json:"earnings" db:"earnings"`
}

// ReferralRepository defines operations for managing referrals
type ReferralRepository interface {
	// GetReferralStats retrieves referral statistics for a user
	GetReferralStats(ctx context.Context, userID string, tenantID string) (ReferralStats, error)
	
	// GetRecentReferrals retrieves recent referrals for a user
	GetRecentReferrals(ctx context.Context, userID string, tenantID string, limit int) ([]ReferralRecord, error)
}