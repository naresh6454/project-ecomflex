package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/naresh6454/ecomflex-backend/internal/domain/entity"
	"github.com/naresh6454/ecomflex-backend/internal/domain/repository"
)

// PostgresUserRepository implements UserRepository interface using PostgreSQL
type PostgresUserRepository struct {
	db *sqlx.DB
}

// NewPostgresUserRepository creates a new PostgresUserRepository
func NewPostgresUserRepository(db *sqlx.DB) repository.UserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}

// CreateUser creates a new user
func (r *PostgresUserRepository) CreateUser(ctx context.Context, user *entity.User) error {
	query := `
		INSERT INTO users (
			id, tenant_id, email, password_hash, full_name, role, phone, 
			referral_code, status, profile_picture, social_links, follower_count, 
			created_at, updated_at
		) VALUES (
			:id, :tenant_id, :email, :password_hash, :full_name, :role, :phone, 
			:referral_code, :status, :profile_picture, :social_links, :follower_count, 
			:created_at, :updated_at
		)
	`
	
	// Create a map for the query parameters
	params := map[string]interface{}{
		"id":              user.ID,
		"tenant_id":       user.TenantID,
		"email":           user.Email,
		"password_hash":   user.PasswordHash,
		"full_name":       user.FullName,
		"role":            user.Role,
		"phone":           user.Phone,
		"referral_code":   user.ReferralCode,
		"status":          user.Status,
		"profile_picture": user.ProfilePicture,
		"social_links":    pq.Array(user.SocialLinks),
		"follower_count":  user.FollowerCount,
		"created_at":      user.CreatedAt,
		"updated_at":      user.UpdatedAt,
	}
	
	// Execute the query
	_, err := r.db.NamedExecContext(ctx, query, params)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			// Check for unique violation error
			if pqErr.Code == "23505" {
				if pqErr.Constraint == "users_email_key" {
					return errors.New("email already exists")
				}
				if pqErr.Constraint == "users_referral_code_key" {
					return errors.New("referral code already exists")
				}
			}
		}
		return fmt.Errorf("failed to create user: %w", err)
	}
	
	return nil
}

// GetUserByID gets a user by ID
func (r *PostgresUserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	query := `
		SELECT 
			id, tenant_id, email, password_hash, full_name, role, phone, 
			referral_code, status, profile_picture, social_links, follower_count, 
			created_at, updated_at, deleted_at 
		FROM users 
		WHERE id = $1 AND deleted_at IS NULL
	`
	
	user := &entity.User{}
	var socialLinks pq.StringArray
	
	// Execute the query
	err := r.db.QueryRowxContext(ctx, query, id).Scan(
		&user.ID, &user.TenantID, &user.Email, &user.PasswordHash, &user.FullName, &user.Role, 
		&user.Phone, &user.ReferralCode, &user.Status, &user.ProfilePicture, &socialLinks, 
		&user.FollowerCount, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}
	
	// Convert pq.StringArray to []string
	user.SocialLinks = []string(socialLinks)
	
	return user, nil
}

// GetUserByEmail gets a user by email
func (r *PostgresUserRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	query := `
		SELECT 
			id, tenant_id, email, password_hash, full_name, role, phone, 
			referral_code, status, profile_picture, social_links, follower_count, 
			created_at, updated_at, deleted_at 
		FROM users 
		WHERE email = $1 AND deleted_at IS NULL
	`
	
	user := &entity.User{}
	var socialLinks pq.StringArray
	
	// Execute the query
	err := r.db.QueryRowxContext(ctx, query, email).Scan(
		&user.ID, &user.TenantID, &user.Email, &user.PasswordHash, &user.FullName, &user.Role, 
		&user.Phone, &user.ReferralCode, &user.Status, &user.ProfilePicture, &socialLinks, 
		&user.FollowerCount, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}
	
	// Convert pq.StringArray to []string
	user.SocialLinks = []string(socialLinks)
	
	return user, nil
}

// GetUserByReferralCode gets a user by referral code
func (r *PostgresUserRepository) GetUserByReferralCode(ctx context.Context, referralCode string) (*entity.User, error) {
	query := `
		SELECT 
			id, tenant_id, email, password_hash, full_name, role, phone, 
			referral_code, status, profile_picture, social_links, follower_count, 
			created_at, updated_at, deleted_at 
		FROM users 
		WHERE referral_code = $1 AND deleted_at IS NULL AND status = 'active'
	`
	
	user := &entity.User{}
	var socialLinks pq.StringArray
	
	// Execute the query
	err := r.db.QueryRowxContext(ctx, query, referralCode).Scan(
		&user.ID, &user.TenantID, &user.Email, &user.PasswordHash, &user.FullName, &user.Role, 
		&user.Phone, &user.ReferralCode, &user.Status, &user.ProfilePicture, &socialLinks, 
		&user.FollowerCount, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("failed to get user by referral code: %w", err)
	}
	
	// Convert pq.StringArray to []string
	user.SocialLinks = []string(socialLinks)
	
	return user, nil
}

// UpdateUser updates a user
func (r *PostgresUserRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	query := `
		UPDATE users 
		SET 
			full_name = :full_name,
			phone = :phone,
			status = :status,
			profile_picture = :profile_picture,
			social_links = :social_links,
			follower_count = :follower_count,
			updated_at = :updated_at
		WHERE id = :id AND deleted_at IS NULL
	`
	
	// Set updated_at
	user.UpdatedAt = time.Now()
	
	// Create a map for the query parameters
	params := map[string]interface{}{
		"id":              user.ID,
		"full_name":       user.FullName,
		"phone":           user.Phone,
		"status":          user.Status,
		"profile_picture": user.ProfilePicture,
		"social_links":    pq.Array(user.SocialLinks),
		"follower_count":  user.FollowerCount,
		"updated_at":      user.UpdatedAt,
	}
	
	// Execute the query
	result, err := r.db.NamedExecContext(ctx, query, params)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	
	// Check if the user was updated
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	
	if rowsAffected == 0 {
		return errors.New("user not found")
	}
	
	return nil
}

// DeleteUser deletes a user (soft delete)
func (r *PostgresUserRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	query := `
		UPDATE users 
		SET 
			deleted_at = $1,
			updated_at = $1
		WHERE id = $2 AND deleted_at IS NULL
	`
	
	now := time.Now()
	
	// Execute the query
	result, err := r.db.ExecContext(ctx, query, now, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	
	// Check if the user was deleted
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	
	if rowsAffected == 0 {
		return errors.New("user not found")
	}
	
	return nil
}

// GetInfluencersByStatus gets influencers by status
func (r *PostgresUserRepository) GetInfluencersByStatus(ctx context.Context, status string, limit, offset int) ([]*entity.User, int, error) {
	// Query to get influencers with the given status
	query := `
		SELECT 
			id, tenant_id, email, password_hash, full_name, role, phone, 
			referral_code, status, profile_picture, social_links, follower_count, 
			created_at, updated_at, deleted_at 
		FROM users 
		WHERE role = 'influencer' AND status = $1 AND deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`
	
	// Query to count total influencers with the given status
	countQuery := `
		SELECT COUNT(*) 
		FROM users 
		WHERE role = 'influencer' AND status = $1 AND deleted_at IS NULL
	`
	
	// Execute count query
	var total int
	err := r.db.GetContext(ctx, &total, countQuery, status)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count influencers: %w", err)
	}
	
	// Execute main query
	rows, err := r.db.QueryxContext(ctx, query, status, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get influencers: %w", err)
	}
	defer rows.Close()
	
	// Process results
	influencers := []*entity.User{}
	for rows.Next() {
		user := &entity.User{}
		var socialLinks pq.StringArray
		
		err := rows.Scan(
			&user.ID, &user.TenantID, &user.Email, &user.PasswordHash, &user.FullName, &user.Role, 
			&user.Phone, &user.ReferralCode, &user.Status, &user.ProfilePicture, &socialLinks, 
			&user.FollowerCount, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt,
		)
		
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan influencer: %w", err)
		}
		
		// Convert pq.StringArray to []string
		user.SocialLinks = []string(socialLinks)
		
		influencers = append(influencers, user)
	}
	
	return influencers, total, nil
}

// IsEmailTaken checks if an email is already taken
func (r *PostgresUserRepository) IsEmailTaken(ctx context.Context, email string) (bool, error) {
	query := `
		SELECT EXISTS (
			SELECT 1 
			FROM users 
			WHERE email = $1 AND deleted_at IS NULL
		)
	`
	
	var exists bool
	err := r.db.GetContext(ctx, &exists, query, email)
	if err != nil {
		return false, fmt.Errorf("failed to check if email is taken: %w", err)
	}
	
	return exists, nil
}

// IsReferralCodeTaken checks if a referral code is already taken
func (r *PostgresUserRepository) IsReferralCodeTaken(ctx context.Context, referralCode string) (bool, error) {
	query := `
		SELECT EXISTS (
			SELECT 1 
			FROM users 
			WHERE referral_code = $1 AND deleted_at IS NULL
		)
	`
	
	var exists bool
	err := r.db.GetContext(ctx, &exists, query, referralCode)
	if err != nil {
		return false, fmt.Errorf("failed to check if referral code is taken: %w", err)
	}
	
	return exists, nil
}