// ecomflex-backend/internal/domain/repository/product_repository.go
package repository

import (
	"context"

	"github.com/naresh6454/ecomflex-backend/internal/domain/entity"
)

type ProductRepository interface {
	Create(ctx context.Context, product *entity.Product) error
	FindByID(ctx context.Context, id string) (*entity.Product, error)
	Update(ctx context.Context, product *entity.Product) error
	Delete(ctx context.Context, id string) error
	FindAll(ctx context.Context, filter entity.ProductFilter) ([]*entity.Product, int64, error)
	UpdateStatus(ctx context.Context, id string, isActive bool) error
	FindTrending(ctx context.Context, limit int) ([]*entity.Product, error)
	GetCategories(ctx context.Context) ([]string, error)
	UpdateImage(ctx context.Context, id string, imageURL string) error
}