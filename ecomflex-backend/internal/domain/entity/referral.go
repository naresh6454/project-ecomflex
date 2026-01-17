package entity

import (
	"time"
	
	"github.com/google/uuid"
)

// Referral represents a referral made by an influencer
type Referral struct {
	ID           uuid.UUID  `json:"id" db:"id"`
	TenantID     uuid.UUID  `json:"tenant_id" db:"tenant_id"`
	InfluencerID uuid.UUID  `json:"influencer_id" db:"influencer_id"`
	UserID       uuid.UUID  `json:"user_id" db:"user_id"`
	ProductID    uuid.UUID  `json:"product_id" db:"product_id"`
	Status       string     `json:"status" db:"status"`
	Earnings     float64    `json:"earnings" db:"earnings"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

// NewReferral creates a new referral
func NewReferral(tenantID, influencerID, userID, productID uuid.UUID, earnings float64) *Referral {
	now := time.Now()
	return &Referral{
		ID:           uuid.New(),
		TenantID:     tenantID,
		InfluencerID: influencerID,
		UserID:       userID,
		ProductID:    productID,
		Status:       "pending",
		Earnings:     earnings,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}

// Approve approves a referral
func (r *Referral) Approve() {
	r.Status = "approved"
	r.UpdatedAt = time.Now()
}

// Reject rejects a referral
func (r *Referral) Reject() {
	r.Status = "rejected"
	r.UpdatedAt = time.Now()
}

// IsPending checks if the referral is pending
func (r *Referral) IsPending() bool {
	return r.Status == "pending"
}

// IsApproved checks if the referral is approved
func (r *Referral) IsApproved() bool {
	return r.Status == "approved"
}

// IsRejected checks if the referral is rejected
func (r *Referral) IsRejected() bool {
	return r.Status == "rejected"
}