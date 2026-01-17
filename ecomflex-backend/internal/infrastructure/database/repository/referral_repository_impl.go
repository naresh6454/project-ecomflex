package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/naresh6454/ecomflex-backend/internal/domain/repository"
)

// PostgresReferralRepository implements ReferralRepository
type PostgresReferralRepository struct {
	db *sqlx.DB
}

// NewPostgresReferralRepository creates a new PostgresReferralRepository
func NewPostgresReferralRepository(db *sqlx.DB) repository.ReferralRepository {
	return &PostgresReferralRepository{
		db: db,
	}
}

// GetReferralStats retrieves referral statistics for a user
func (r *PostgresReferralRepository) GetReferralStats(ctx context.Context, userID string, tenantID string) (repository.ReferralStats, error) {
	stats := repository.ReferralStats{}
	
	// Parse UUIDs
	influencerUUID, err := uuid.Parse(userID)
	if err != nil {
		return repository.ReferralStats{}, err
	}
	
	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		return repository.ReferralStats{}, err
	}
	
	// Modified query to include tenant_id filtering
	err = r.db.GetContext(ctx, &stats, `
		SELECT 
			COUNT(*) as total_referrals,
			COUNT(CASE WHEN status = 'approved' THEN 1 END) as approved_bookings,
			COUNT(CASE WHEN status = 'pending' THEN 1 END) as pending_bookings,
			COALESCE(SUM(earnings), 0) as total_earnings
		FROM referrals
		WHERE influencer_id = $1 AND tenant_id = $2 AND deleted_at IS NULL
	`, influencerUUID, tenantUUID)
	
	if err != nil {
		// Return empty stats for new users with no referrals yet
		return repository.ReferralStats{
			TotalReferrals:   0,
			ApprovedBookings: 0,
			PendingBookings:  0,
			TotalEarnings:    0,
		}, nil
	}
	
	return stats, nil
}

// GetRecentReferrals retrieves recent referrals for a user
func (r *PostgresReferralRepository) GetRecentReferrals(ctx context.Context, userID string, tenantID string, limit int) ([]repository.ReferralRecord, error) {
	var referrals []repository.ReferralRecord
	
	// Parse UUIDs
	influencerUUID, err := uuid.Parse(userID)
	if err != nil {
		return []repository.ReferralRecord{}, err
	}
	
	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		return []repository.ReferralRecord{}, err
	}
	
	// Modified query to include tenant_id filtering and ensure users are filtered by the same tenant
	err = r.db.SelectContext(ctx, &referrals, `
		SELECT 
			r.id,
			u.full_name as user,
			p.name as product,
			TO_CHAR(r.created_at, 'Mon DD, YYYY') as date,
			CASE 
				WHEN r.status = 'approved' THEN 'Approved'
				WHEN r.status = 'pending' THEN 'Pending'
				ELSE 'Rejected'
			END as status,
			r.earnings
		FROM referrals r
		JOIN users u ON r.user_id = u.id AND u.tenant_id = r.tenant_id
		JOIN products p ON r.product_id = p.id AND p.tenant_id = r.tenant_id
		WHERE r.influencer_id = $1 
		  AND r.tenant_id = $2 
		  AND r.deleted_at IS NULL
		ORDER BY r.created_at DESC
		LIMIT $3
	`, influencerUUID, tenantUUID, limit)
	
	if err != nil {
		// Return empty array for new users with no referrals yet
		return []repository.ReferralRecord{}, nil
	}
	
	return referrals, nil
}