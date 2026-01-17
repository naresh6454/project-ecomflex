package repository

import (
	"context"
	"database/sql"
	"encoding/json" // Add this import
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/naresh6454/ecomflex-backend/internal/domain/entity"
	"github.com/naresh6454/ecomflex-backend/internal/domain/repository"
)

// TenantRepositoryImpl implements TenantRepository interface
type TenantRepositoryImpl struct {
	db *sqlx.DB
}

// NewTenantRepository creates a new TenantRepositoryImpl
func NewTenantRepository(db *sqlx.DB) repository.TenantRepository {
	return &TenantRepositoryImpl{
		db: db,
	}
}

// CreateTenant creates a new tenant
func (r *TenantRepositoryImpl) CreateTenant(ctx context.Context, tenant *entity.Tenant) error {
	// Convert map to JSON before storing
	settingsJSON, err := json.Marshal(tenant.Settings)
	if err != nil {
		return fmt.Errorf("failed to marshal tenant settings: %w", err)
	}

	query := `
		INSERT INTO tenants (
			id, name, domain, plan_id, is_active, settings, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8
		)
	`

	_, err = r.db.ExecContext(
		ctx,
		query,
		tenant.ID,
		tenant.Name,
		tenant.Domain,
		tenant.PlanID,
		tenant.IsActive,
		settingsJSON, // Pass the JSON representation instead of the map
		tenant.CreatedAt,
		tenant.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create tenant: %w", err)
	}

	return nil
}

// GetTenantByID retrieves a tenant by ID
func (r *TenantRepositoryImpl) GetTenantByID(ctx context.Context, id uuid.UUID) (*entity.Tenant, error) {
	query := `
		SELECT id, name, domain, plan_id, is_active, settings, created_at, updated_at, deleted_at
		FROM tenants
		WHERE id = $1 AND deleted_at IS NULL
	`

	var tenant entity.Tenant
	var settingsJSON []byte
	
	err := r.db.QueryRowxContext(ctx, query, id).Scan(
		&tenant.ID,
		&tenant.Name,
		&tenant.Domain,
		&tenant.PlanID,
		&tenant.IsActive,
		&settingsJSON,
		&tenant.CreatedAt,
		&tenant.UpdatedAt,
		&tenant.DeletedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("tenant not found")
		}
		return nil, fmt.Errorf("failed to get tenant: %w", err)
	}
	
	// Unmarshal the settings
	tenant.Settings = make(map[string]interface{})
	if settingsJSON != nil {
		if err := json.Unmarshal(settingsJSON, &tenant.Settings); err != nil {
			return nil, fmt.Errorf("failed to unmarshal tenant settings: %w", err)
		}
	}

	return &tenant, nil
}

// GetTenantByDomain retrieves a tenant by domain
func (r *TenantRepositoryImpl) GetTenantByDomain(ctx context.Context, domain string) (*entity.Tenant, error) {
	query := `
		SELECT id, name, domain, plan_id, is_active, settings, created_at, updated_at, deleted_at
		FROM tenants
		WHERE domain = $1 AND deleted_at IS NULL
	`

	var tenant entity.Tenant
	var settingsJSON []byte
	
	err := r.db.QueryRowxContext(ctx, query, domain).Scan(
		&tenant.ID,
		&tenant.Name,
		&tenant.Domain,
		&tenant.PlanID,
		&tenant.IsActive,
		&settingsJSON,
		&tenant.CreatedAt,
		&tenant.UpdatedAt,
		&tenant.DeletedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("tenant not found")
		}
		return nil, fmt.Errorf("failed to get tenant: %w", err)
	}
	
	// Unmarshal the settings
	tenant.Settings = make(map[string]interface{})
	if settingsJSON != nil {
		if err := json.Unmarshal(settingsJSON, &tenant.Settings); err != nil {
			return nil, fmt.Errorf("failed to unmarshal tenant settings: %w", err)
		}
	}

	return &tenant, nil
}

// UpdateTenant updates a tenant
func (r *TenantRepositoryImpl) UpdateTenant(ctx context.Context, tenant *entity.Tenant) error {
	// Convert map to JSON before storing
	settingsJSON, err := json.Marshal(tenant.Settings)
	if err != nil {
		return fmt.Errorf("failed to marshal tenant settings: %w", err)
	}
	
	query := `
		UPDATE tenants
		SET name = $1, domain = $2, plan_id = $3, is_active = $4,
			settings = $5, updated_at = $6
		WHERE id = $7 AND deleted_at IS NULL
	`

	_, err = r.db.ExecContext(
		ctx,
		query,
		tenant.Name,
		tenant.Domain,
		tenant.PlanID,
		tenant.IsActive,
		settingsJSON, // Pass the JSON representation instead of the map
		tenant.UpdatedAt,
		tenant.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update tenant: %w", err)
	}

	return nil
}

// DeleteTenant marks a tenant as deleted
func (r *TenantRepositoryImpl) DeleteTenant(ctx context.Context, id uuid.UUID) error {
	query := `
		UPDATE tenants
		SET deleted_at = NOW(), updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete tenant: %w", err)
	}

	return nil
}