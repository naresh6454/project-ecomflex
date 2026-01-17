// ecomflex-backend/internal/infrastructure/database/repository/product_repository_impl.go
package repository

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/naresh6454/ecomflex-backend/internal/domain/entity"
	domainRepo "github.com/naresh6454/ecomflex-backend/internal/domain/repository"
)

type productRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) domainRepo.ProductRepository {
	return &productRepositoryImpl{
		db: db,
	}
}

func (r *productRepositoryImpl) Create(ctx context.Context, product *entity.Product) error {
	return r.db.WithContext(ctx).Create(product).Error
}

func (r *productRepositoryImpl) FindByID(ctx context.Context, id string) (*entity.Product, error) {
	var product entity.Product
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("product not found with id: %s", id)
		}
		return nil, err
	}
	return &product, nil
}

func (r *productRepositoryImpl) Update(ctx context.Context, product *entity.Product) error {
	return r.db.WithContext(ctx).Save(product).Error
}

func (r *productRepositoryImpl) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.Product{}).Error
}

func (r *productRepositoryImpl) FindAll(ctx context.Context, filter entity.ProductFilter) ([]*entity.Product, int64, error) {
	var products []*entity.Product
	var total int64

	query := r.db.WithContext(ctx).Model(&entity.Product{})

	// Apply filters
	if filter.Search != "" {
		query = query.Where("name ILIKE ? OR asin ILIKE ? OR founder ILIKE ? OR details ILIKE ?", 
			"%"+filter.Search+"%", "%"+filter.Search+"%", "%"+filter.Search+"%", "%"+filter.Search+"%")
	}
	
	if filter.Category != "" {
		query = query.Where("category = ?", filter.Category)
	}
	
	if filter.IsActive != nil {
		query = query.Where("is_active = ?", *filter.IsActive)
	}

	// Get total count
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination
	if filter.Page > 0 && filter.Limit > 0 {
		offset := (filter.Page - 1) * filter.Limit
		query = query.Offset(offset).Limit(filter.Limit)
	}

	// Apply sorting
	switch filter.SortBy {
	case "newest":
		query = query.Order("created_at DESC")
	case "name":
		query = query.Order("name ASC")
	case "popular":
		query = query.Order("current_bookings DESC")
	default:
		query = query.Order("created_at DESC")
	}

	if err := query.Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

func (r *productRepositoryImpl) UpdateStatus(ctx context.Context, id string, isActive bool) error {
	return r.db.WithContext(ctx).Model(&entity.Product{}).Where("id = ?", id).Update("is_active", isActive).Error
}

func (r *productRepositoryImpl) FindTrending(ctx context.Context, limit int) ([]*entity.Product, error) {
	var products []*entity.Product
	if err := r.db.WithContext(ctx).
		Where("is_active = ?", true).
		Order("current_bookings DESC").
		Limit(limit).
		Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepositoryImpl) GetCategories(ctx context.Context) ([]string, error) {
	var categories []string
	if err := r.db.WithContext(ctx).Model(&entity.Product{}).
		Distinct("category").
		Pluck("category", &categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *productRepositoryImpl) UpdateImage(ctx context.Context, id string, imageURL string) error {
	return r.db.WithContext(ctx).Model(&entity.Product{}).
		Where("id = ?", id).
		Update("image_url", imageURL).Error
}