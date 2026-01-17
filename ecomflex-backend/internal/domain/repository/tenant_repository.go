package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/naresh6454/ecomflex-backend/internal/domain/entity"
)

// TenantRepository defines operations for managing tenants
type TenantRepository interface {
	// CreateTenant creates a new tenant
	CreateTenant(ctx context.Context, tenant *entity.Tenant) error

	// GetTenantByID retrieves a tenant by ID
	GetTenantByID(ctx context.Context, id uuid.UUID) (*entity.Tenant, error)

	// GetTenantByDomain retrieves a tenant by domain
	GetTenantByDomain(ctx context.Context, domain string) (*entity.Tenant, error)

	// UpdateTenant updates a tenant
	UpdateTenant(ctx context.Context, tenant *entity.Tenant) error

	// DeleteTenant marks a tenant as deleted
	DeleteTenant(ctx context.Context, id uuid.UUID) error
}