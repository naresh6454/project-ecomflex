// ecomflex-backend/internal/domain/entity/product.go
package entity

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID               uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name             string    `json:"name" gorm:"not null"`
	ASIN             string    `json:"asin" gorm:"not null"`
	Founder          string    `json:"founder" gorm:"not null"`
	ProductLink      string    `json:"productLink" gorm:"not null"`
	Details          string    `json:"details" gorm:"type:text"`
	RequiredBookings int       `json:"requiredBookings" gorm:"not null;default:0"`
	CurrentBookings  int       `json:"currentBookings" gorm:"not null;default:0"`
	ImageURL         string    `json:"image" gorm:"column:image_url"`
	IsActive         bool      `json:"isActive" gorm:"not null;default:true"`
	TenantID         uuid.UUID `json:"tenantId" gorm:"type:uuid;not null"`
	CreatedAt        time.Time `json:"createdAt" gorm:"not null"`
	UpdatedAt        time.Time `json:"updatedAt" gorm:"not null"`
}

type ProductFilter struct {
	Category string
	Search   string
	SortBy   string
	Page     int
	Limit    int
	IsActive *bool
}
