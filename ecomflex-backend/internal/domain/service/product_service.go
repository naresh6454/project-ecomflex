// ecomflex-backend/internal/domain/service/product_service.go
package service

import (
	"context"
	"mime/multipart"

	"github.com/naresh6454/ecomflex-backend/internal/domain/entity"
)

type ProductService interface {
	CreateProduct(ctx context.Context, product *entity.Product) error
	GetProductByID(ctx context.Context, id string) (*entity.Product, error)
	UpdateProduct(ctx context.Context, product *entity.Product) error
	DeleteProduct(ctx context.Context, id string) error
	GetProducts(ctx context.Context, filter entity.ProductFilter) ([]*entity.Product, int64, error)
	ToggleProductStatus(ctx context.Context, id string, isActive bool) error
	GetTrendingProducts(ctx context.Context, limit int) ([]*entity.Product, error)
	GetProductCategories(ctx context.Context) ([]string, error)
	UploadProductImage(ctx context.Context, id string, file *multipart.FileHeader) (string, error)
}