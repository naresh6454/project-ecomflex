// ecomflex-backend/internal/app/service/product_service_impl.go
package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/naresh6454/ecomflex-backend/internal/domain/entity"
	"github.com/naresh6454/ecomflex-backend/internal/domain/repository"
	"github.com/naresh6454/ecomflex-backend/internal/domain/service"
)

type productServiceImpl struct {
	productRepository repository.ProductRepository
	// Could add S3Client or other file storage service here
}

func NewProductService(productRepo repository.ProductRepository) service.ProductService {
	return &productServiceImpl{
		productRepository: productRepo,
	}
}

func (s *productServiceImpl) CreateProduct(ctx context.Context, product *entity.Product) error {
	// Set creation time
	now := time.Now()
	product.CreatedAt = now
	product.UpdatedAt = now
	product.CurrentBookings = 0

	return s.productRepository.Create(ctx, product)
}

func (s *productServiceImpl) GetProductByID(ctx context.Context, id string) (*entity.Product, error) {
	return s.productRepository.FindByID(ctx, id)
}

// Helper function to convert uuid.UUID to string
func (s *productServiceImpl) uuidToString(id interface{}) string {
	if id == nil {
		return ""
	}
	switch v := id.(type) {
	case string:
		return v
	default:
		return fmt.Sprint(v)
	}
}

func (s *productServiceImpl) UpdateProduct(ctx context.Context, product *entity.Product) error {
	// Verify the product exists
	existingProduct, err := s.productRepository.FindByID(ctx, product.ID.String())
	if err != nil {
		return err
	}

	// Update timestamps and preserve some fields
	product.CreatedAt = existingProduct.CreatedAt
	product.UpdatedAt = time.Now()
	product.CurrentBookings = existingProduct.CurrentBookings

	// If image was not provided, preserve the existing one
	if product.ImageURL == "" {
		product.ImageURL = existingProduct.ImageURL
	}

	return s.productRepository.Update(ctx, product)
}

func (s *productServiceImpl) DeleteProduct(ctx context.Context, id string) error {
	// Verify the product exists
	_, err := s.productRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}
	return s.productRepository.Delete(ctx, id)
}

func (s *productServiceImpl) GetProducts(ctx context.Context, filter entity.ProductFilter) ([]*entity.Product, int64, error) {
	return s.productRepository.FindAll(ctx, filter)
}

func (s *productServiceImpl) ToggleProductStatus(ctx context.Context, id string, isActive bool) error {
	// Verify the product exists
	_, err := s.productRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}
	return s.productRepository.UpdateStatus(ctx, id, isActive)
}

func (s *productServiceImpl) GetTrendingProducts(ctx context.Context, limit int) ([]*entity.Product, error) {
	return s.productRepository.FindTrending(ctx, limit)
}

func (s *productServiceImpl) GetProductCategories(ctx context.Context) ([]string, error) {
	return s.productRepository.GetCategories(ctx)
}

// ecomflex-backend/internal/app/service/product_service_impl.go (continued)

func (s *productServiceImpl) UploadProductImage(ctx context.Context, id string, file *multipart.FileHeader) (string, error) {
	// Verify the product exists
	_, err := s.productRepository.FindByID(ctx, id)
	if err != nil {
		return "", err
	}

	// Generate a unique filename
	filename := fmt.Sprintf("%s-%d%s", id, time.Now().UnixNano(), filepath.Ext(file.Filename))

	// In a real implementation, you would upload the file to S3 or similar service
	// For this example, we'll assume the image was uploaded and return a URL
	imageURL := fmt.Sprintf("https://your-s3-bucket.s3.amazonaws.com/products/%s", filename)

	// Update the product with the new image URL
	err = s.productRepository.UpdateImage(ctx, id, imageURL)
	if err != nil {
		return "", err
	}

	return imageURL, nil
}
