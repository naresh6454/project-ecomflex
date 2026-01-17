// ecomflex-backend/internal/api/route/product_routes.go
package route

import (
	"github.com/gin-gonic/gin"
	
	"github.com/naresh6454/ecomflex-backend/internal/api/handler"
	"github.com/naresh6454/ecomflex-backend/internal/api/middleware"
)

func ConfigureProductRoutes(
	router *gin.RouterGroup,
	productHandler *handler.ProductHandler,
	authMiddleware *middleware.AuthMiddleware,
) {
	// Public routes - anyone can access
	publicRoutes := router.Group("/products")
	{
		publicRoutes.GET("", productHandler.GetProducts)
		publicRoutes.GET("/:id", productHandler.GetProduct)
		publicRoutes.GET("/trending", productHandler.GetTrendingProducts)
		publicRoutes.GET("/categories", productHandler.GetProductCategories)
	}

	// Admin routes - only super admin can access
	adminRoutes := router.Group("/admin/products")
	// Note: Update the method name based on your AuthMiddleware implementation
	// Commonly used names might be: Authenticate(), Auth(), CheckAuth(), etc.
	adminRoutes.Use(authMiddleware.Authenticate()) // Changed from AuthRequired()
	adminRoutes.Use(authMiddleware.RequireRole("super_admin"))
	{
		adminRoutes.POST("", productHandler.CreateProduct)
		adminRoutes.PUT("/:id", productHandler.UpdateProduct)
		adminRoutes.DELETE("/:id", productHandler.DeleteProduct)
		adminRoutes.PATCH("/:id/status", productHandler.ToggleProductStatus)
		adminRoutes.POST("/:id/image", productHandler.UploadProductImage)
	}
}