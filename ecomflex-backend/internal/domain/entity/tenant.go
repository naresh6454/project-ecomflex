package entity

import (
	"time"

	"github.com/google/uuid"
)

// Tenant represents a tenant in the multi-tenant system
type Tenant struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	Name        string     `json:"name" db:"name"`
	Domain      string     `json:"domain" db:"domain"`
	PlanID      string     `json:"plan_id" db:"plan_id"`
	IsActive    bool       `json:"is_active" db:"is_active"`
	Settings    map[string]interface{} `json:"settings,omitempty" db:"settings"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

// NewTenant creates a new tenant
func NewTenant(name, domain, planID string) *Tenant {
	now := time.Now()
	return &Tenant{
		ID:        uuid.New(),
		Name:      name,
		Domain:    domain,
		PlanID:    planID,
		IsActive:  true,
		Settings:  make(map[string]interface{}),
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// UpdateSettings updates tenant settings
func (t *Tenant) UpdateSettings(settings map[string]interface{}) {
	t.Settings = settings
	t.UpdatedAt = time.Now()
}

// Deactivate deactivates the tenant
func (t *Tenant) Deactivate() {
	t.IsActive = false
	t.UpdatedAt = time.Now()
}

// Activate activates the tenant
func (t *Tenant) Activate() {
	t.IsActive = true
	t.UpdatedAt = time.Now()
}