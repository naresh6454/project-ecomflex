package entity

import (
	"time"

	"github.com/google/uuid"
)

// Role represents user role in the system
type Role string

// User roles
const (
	RoleSuperAdmin Role = "super_admin"
	RoleInfluencer Role = "influencer"
	RolePublic     Role = "public"
)

// User represents a user entity
type User struct {
	ID             uuid.UUID  `json:"id" db:"id"`
	TenantID       uuid.UUID  `json:"tenant_id" db:"tenant_id"`
	Email          string     `json:"email" db:"email"`
	PasswordHash   string     `json:"-" db:"password_hash"`
	FullName       string     `json:"full_name" db:"full_name"`
	Role           Role       `json:"role" db:"role"`
	Phone          string     `json:"phone,omitempty" db:"phone"`
	ReferralCode   string     `json:"referral_code,omitempty" db:"referral_code"`
	Status         string     `json:"status" db:"status"`
	ProfilePicture string     `json:"profile_picture,omitempty" db:"profile_picture"`
	SocialLinks    []string   `json:"social_links,omitempty" db:"social_links"`
	FollowerCount  int        `json:"follower_count,omitempty" db:"follower_count"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

// NewUser creates a new user instance
func NewUser(
	tenantID uuid.UUID,
	email string,
	passwordHash string,
	fullName string,
	role Role,
	phone string,
) *User {
	now := time.Now()
	return &User{
		ID:           uuid.New(),
		TenantID:     tenantID,
		Email:        email,
		PasswordHash: passwordHash,
		FullName:     fullName,
		Role:         role,
		Phone:        phone,
		Status:       "active",
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}

// NewInfluencer creates a new influencer user
func NewInfluencer(
	tenantID uuid.UUID,
	email string,
	passwordHash string,
	fullName string,
	phone string,
	referralCode string,
	socialLinks []string,
	followerCount int,
) *User {
	user := NewUser(tenantID, email, passwordHash, fullName, RoleInfluencer, phone)
	user.ReferralCode = referralCode
	user.SocialLinks = socialLinks
	user.FollowerCount = followerCount
	user.Status = "pending_approval" // Influencers require approval
	return user
}

// IsActive checks if user is active
func (u *User) IsActive() bool {
	return u.Status == "active" && u.DeletedAt == nil
}

// IsPendingApproval checks if user is pending approval
func (u *User) IsPendingApproval() bool {
	return u.Status == "pending_approval"
}

// SetActive sets user status to active
func (u *User) SetActive() {
	u.Status = "active"
	u.UpdatedAt = time.Now()
}

// SetInactive sets user status to inactive
func (u *User) SetInactive() {
	u.Status = "inactive"
	u.UpdatedAt = time.Now()
}

// SetRejected sets user status to rejected (for influencers)
func (u *User) SetRejected() {
	u.Status = "rejected"
	u.UpdatedAt = time.Now()
}
