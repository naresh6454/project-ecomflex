// ecomflex-backend/internal/api/handler/product_handler.go
package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/naresh6454/ecomflex-backend/internal/api/response"
	"github.com/naresh6454/ecomflex-backend/internal/domain/entity"
	"github.com/naresh6454/ecomflex-backend/internal/domain/service"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

// ProductRequest defines the request for creating/updating products
type ProductRequest struct {
	Name             string `json:"name" binding:"required"`
	ASIN             string `json:"asin" binding:"required"`
	Founder          string `json:"founder" binding:"required"`
	ProductLink      string `json:"productLink" binding:"required"`
	Details          string `json:"details"`
	RequiredBookings int    `json:"requiredBookings" binding:"required"`
	IsActive         bool   `json:"isActive"`
}

// CreateProduct handles product creation
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	// Get tenant ID from context (set by middleware)
	tenantID, exists := c.Get("tenantID")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "Unauthorized: missing tenant ID", nil)
		return
	}

	var req ProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request: "+err.Error(), nil)
		return
	}

	// Convert tenant ID string to UUID
	tenantUUID, err := uuid.Parse(tenantID.(string))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid tenant ID format", nil)
		return
	}

	product := &entity.Product{
		Name:             req.Name,
		ASIN:             req.ASIN,
		Founder:          req.Founder,
		ProductLink:      req.ProductLink,
		Details:          req.Details,
		RequiredBookings: req.RequiredBookings,
		IsActive:         req.IsActive,
		TenantID:         tenantUUID,
	}

	if err := h.productService.CreateProduct(c, product); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create product: "+err.Error(), nil)
		return
	}

	response.Success(c, http.StatusCreated, "Product created successfully", product)
}

// GetProduct handles retrieving a single product
func (h *ProductHandler) GetProduct(c *gin.Context) {
	id := c.Param("id")
	product, err := h.productService.GetProductByID(c, id)
	if err != nil {
		response.Error(c, http.StatusNotFound, "Product not found: "+err.Error(), nil)
		return
	}

	response.Success(c, http.StatusOK, "Product retrieved successfully", product)
}

// UpdateProduct handles product updates
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	var req ProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request: "+err.Error(), nil)
		return
	}

	// Get tenant ID from context (set by middleware)
	tenantID, exists := c.Get("tenantID")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "Unauthorized: missing tenant ID", nil)
		return
	}

	// Convert ID string to UUID
	productUUID, err := uuid.Parse(id)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid product ID format", nil)
		return
	}

	// Convert tenant ID string to UUID
	tenantUUID, err := uuid.Parse(tenantID.(string))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid tenant ID format", nil)
		return
	}

	product := &entity.Product{
		ID:               productUUID,
		Name:             req.Name,
		ASIN:             req.ASIN,
		Founder:          req.Founder,
		ProductLink:      req.ProductLink,
		Details:          req.Details,
		RequiredBookings: req.RequiredBookings,
		IsActive:         req.IsActive,
		TenantID:         tenantUUID,
	}

	if err := h.productService.UpdateProduct(c, product); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to update product: "+err.Error(), nil)
		return
	}

	response.Success(c, http.StatusOK, "Product updated successfully", product)
}

// DeleteProduct handles product deletion
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := h.productService.DeleteProduct(c, id); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to delete product: "+err.Error(), nil)
		return
	}

	response.Success(c, http.StatusOK, "Product deleted successfully", nil)
}

// GetProducts handles retrieving a list of products with pagination
func (h *ProductHandler) GetProducts(c *gin.Context) {
	// Extract filter parameters
	filter := entity.ProductFilter{
		Search:   c.Query("search"),
		Category: c.Query("category"),
		SortBy:   c.Query("sortBy"),
	}

	// Parse pagination parameters
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	filter.Page = page

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}
	filter.Limit = limit

	// Parse active filter if provided
	if activeStr := c.Query("active"); activeStr != "" {
		active := activeStr == "true"
		filter.IsActive = &active
	}

	products, total, err := h.productService.GetProducts(c, filter)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get products: "+err.Error(), nil)
		return
	}

	// Prepare pagination metadata
	meta := map[string]interface{}{
		"total":       total,
		"page":        filter.Page,
		"limit":       filter.Limit,
		"total_pages": (total + int64(filter.Limit) - 1) / int64(filter.Limit),
	}

	response.Success(c, http.StatusOK, "Products retrieved successfully", gin.H{
		"products": products,
		"meta":     meta,
	})
}

// ToggleProductStatus handles activating/deactivating a product
func (h *ProductHandler) ToggleProductStatus(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		IsActive bool `json:"isActive"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request: "+err.Error(), nil)
		return
	}

	if err := h.productService.ToggleProductStatus(c, id, req.IsActive); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to update product status: "+err.Error(), nil)
		return
	}

	response.Success(c, http.StatusOK, "Product status updated successfully", nil)
}

// UploadProductImage handles product image uploads
func (h *ProductHandler) UploadProductImage(c *gin.Context) {
	id := c.Param("id")

	// Get the file from the request
	file, err := c.FormFile("image")
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid file: "+err.Error(), nil)
		return
	}

	// Upload the file
	imageURL, err := h.productService.UploadProductImage(c, id, file)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to upload image: "+err.Error(), nil)
		return
	}

	response.Success(c, http.StatusOK, "Image uploaded successfully", gin.H{
		"imageUrl": imageURL,
	})
}

// GetTrendingProducts handles retrieving trending products
func (h *ProductHandler) GetTrendingProducts(c *gin.Context) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "5"))
	if err != nil || limit < 1 {
		limit = 5
	}

	products, err := h.productService.GetTrendingProducts(c, limit)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get trending products: "+err.Error(), nil)
		return
	}

	response.Success(c, http.StatusOK, "Trending products retrieved successfully", products)
}

// GetProductCategories handles retrieving product categories
func (h *ProductHandler) GetProductCategories(c *gin.Context) {
	categories, err := h.productService.GetProductCategories(c)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get product categories: "+err.Error(), nil)
		return
	}

	response.Success(c, http.StatusOK, "Product categories retrieved successfully", categories)
}
