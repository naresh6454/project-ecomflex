package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/naresh6454/ecomflex-backend/config"
	"github.com/naresh6454/ecomflex-backend/internal/api/handler"
	"github.com/naresh6454/ecomflex-backend/internal/api/route"
	"github.com/naresh6454/ecomflex-backend/internal/app/service"
	"github.com/naresh6454/ecomflex-backend/internal/domain/entity"
	"github.com/naresh6454/ecomflex-backend/internal/infrastructure/cache"
	"github.com/naresh6454/ecomflex-backend/internal/infrastructure/database"
	repo "github.com/naresh6454/ecomflex-backend/internal/infrastructure/database/repository"
	loggerPkg "github.com/naresh6454/ecomflex-backend/internal/util/logger"
	"github.com/naresh6454/ecomflex-backend/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load configurations
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Check if migration command is provided
	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		log.Println("Running database migrations...")
		if err := migrations.RunMigrations(cfg); err != nil {
			log.Fatalf("Failed to run migrations: %v", err)
		}
		log.Println("Database migrations completed successfully")
		return
	}

	// Initialize logger
	logger := loggerPkg.NewLogger(cfg.LogLevel)
	logger.Info("Starting Ecomflex API server...")

	// Connect to PostgreSQL using sqlx
	sqlxDB, err := database.NewPostgresConnection(cfg.Database)
	if err != nil {
		logger.Fatal("Failed to connect to database", err)
	}
	defer sqlxDB.Close()

	// Initialize a GORM DB connection for product repository
	dsn := os.Getenv("DATABASE_URL") // Try environment variable first
	if dsn == "" {
		// Fallback to constructing from your config
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cfg.Database.Host, cfg.Database.Port, cfg.Database.Username,
			cfg.Database.Password, cfg.Database.DBName)
	}

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("Failed to connect to database with GORM", err)
	}

	// Enable UUID extension first
	if err := gormDB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error; err != nil {
		logger.Fatal("Failed to create UUID extension", err)
	}

	// Auto Migrate the product table
	err = gormDB.AutoMigrate(&entity.Product{})
	if err != nil {
		logger.Fatal("Failed to migrate product table", err)
	}

	// Connect to Redis
	redisClient, err := cache.NewRedisClient(cfg.Redis)
	if err != nil {
		logger.Fatal("Failed to connect to Redis", err)
	}
	defer redisClient.Close()

	// Get AWS S3 credentials from environment variables
	accessKey := os.Getenv("AKIARdbhsbP5LGTLHWW")
	secretKey := os.Getenv("6EvGUth6Dwebfjbfbj7jYAtfbhrbUMU0LtfApO5atqpGN3")
	region := os.Getenv("north-1")
	bucketName := os.Getenv("ecomflex-uploads-dev")

	// If environment variables aren't set, use default values for development
	if accessKey == "" || secretKey == "" {
		logger.Warn("AWS credentials not found in environment variables. Using defaults for development.")
		// For security, we won't hardcode credentials anymore
		logger.Warn("Please set AWS_ACCESS_KEY and AWS_SECRET_KEY environment variables")
	}

	if region == "" {
		region = "eu-north-1" // Stockholm region
	}

	if bucketName == "" {
		bucketName = "ecomflex-uploads-dev"
	}

	// Log S3 configuration for debugging
	logger.Info(fmt.Sprintf("S3 Configuration - Bucket: %s, Region: %s", bucketName, region))
	logger.Info("AWS credentials loaded but S3 client is not yet implemented")
	logger.Info("To implement S3 storage, create the required files in the storage package")

	// Initialize router
	router := route.SetupRouter(cfg, logger, sqlxDB, redisClient)

	// Initialize product repository and service
	productRepository := repo.NewProductRepository(gormDB)
	productService := service.NewProductService(productRepository)

	// Initialize product handler
	productHandler := handler.NewProductHandler(productService)

	// Set up API routes
	apiGroup := router.Group("/api")

	// Set up all product routes in one group (for backward compatibility)
	productsGroup := apiGroup.Group("/products")
	{
		// Public GET routes
		productsGroup.GET("", productHandler.GetProducts)
		productsGroup.GET("/:id", productHandler.GetProduct)
		productsGroup.GET("/trending", productHandler.GetTrendingProducts)
		productsGroup.GET("/categories", productHandler.GetProductCategories)

		// Admin routes - added directly to the same group without auth middleware
		// For development, add a simple middleware that sets tenant ID
		productsGroup.POST("", addDevTenant, productHandler.CreateProduct)
		productsGroup.PUT("/:id", addDevTenant, productHandler.UpdateProduct)
		productsGroup.DELETE("/:id", addDevTenant, productHandler.DeleteProduct)
		productsGroup.PATCH("/:id/status", addDevTenant, productHandler.ToggleProductStatus)
		productsGroup.POST("/:id/image", addDevTenant, productHandler.UploadProductImage)
	}

	// Set up versioned product routes under /api/v1/products
	v1Group := apiGroup.Group("/v1")
	v1ProductsGroup := v1Group.Group("/products")
	{
		// Public GET routes
		v1ProductsGroup.GET("", productHandler.GetProducts)
		v1ProductsGroup.GET("/:id", productHandler.GetProduct)
		v1ProductsGroup.GET("/trending", productHandler.GetTrendingProducts)
		v1ProductsGroup.GET("/categories", productHandler.GetProductCategories)

		// Admin routes
		v1ProductsGroup.POST("", addDevTenant, productHandler.CreateProduct)
		v1ProductsGroup.PUT("/:id", addDevTenant, productHandler.UpdateProduct)
		v1ProductsGroup.DELETE("/:id", addDevTenant, productHandler.DeleteProduct)
		v1ProductsGroup.PATCH("/:id/status", addDevTenant, productHandler.ToggleProductStatus)
		v1ProductsGroup.POST("/:id/image", addDevTenant, productHandler.UploadProductImage)
	}

	// Determine port: environment variable overrides config
	port := cfg.Server.Port
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	// Create HTTP server
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	// Start server in a goroutine
	go func() {
		logger.Info("Server is running on port " + port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Failed to start server", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown the server
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown", err)
	}

	logger.Info("Server exited properly")
}

// addDevTenant is a middleware that adds a development tenant ID to the context
// This is only for development purposes to bypass authentication
func addDevTenant(c *gin.Context) {
	// Add a proper UUID for tenant ID to match the expected format
	// This is important for database operations that expect a valid UUID
	defaultTenantID := "00000000-0000-0000-0000-000000000000"
	tenantUUID, err := uuid.Parse(defaultTenantID)
	if err != nil {
		// Should never happen with a valid UUID string
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Invalid tenant ID format"})
		return
	}

	// Set properly formatted tenant ID in context
	c.Set("tenantID", tenantUUID.String())

	// Set other context values for development
	c.Set("userID", uuid.New().String()) // Generate a random UUID for user ID
	c.Set("role", "super_admin")
	c.Next()
}
