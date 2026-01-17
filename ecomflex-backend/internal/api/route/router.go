package route

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/naresh6454/ecomflex-backend/config"
	"github.com/naresh6454/ecomflex-backend/internal/api/handler"
	"github.com/naresh6454/ecomflex-backend/internal/api/middleware"
	"github.com/naresh6454/ecomflex-backend/internal/app/service"
	"github.com/naresh6454/ecomflex-backend/internal/infrastructure/auth"
	"github.com/naresh6454/ecomflex-backend/internal/infrastructure/cache"
	dbRepo "github.com/naresh6454/ecomflex-backend/internal/infrastructure/database/repository"
	loggerPkg "github.com/naresh6454/ecomflex-backend/internal/util/logger"
)

// SetupRouter sets up the router with all routes and middleware
func SetupRouter(cfg *config.Config, logger loggerPkg.Logger, db *sqlx.DB, redisClient *cache.RedisClient) *gin.Engine {
	// Create router
	router := gin.New()
	
	// Set release mode in production
	if gin.Mode() == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	
	// Add middleware
	router.Use(gin.Recovery())
	router.Use(middleware.CorsMiddleware(&cfg.Cors))
	router.Use(middleware.RequestLogger(logger))
	
	// Create repositories
	userRepo := dbRepo.NewPostgresUserRepository(db)
	authRepo := dbRepo.NewRedisAuthRepository(redisClient.Client)
	tenantRepo := dbRepo.NewTenantRepository(db) // Add tenant repository
	
	// Create a referral repository for influencer functionality
	referralRepo := dbRepo.NewPostgresReferralRepository(db)
	
	// Create JWT provider
	jwtProvider := auth.NewJWTProvider(&cfg.JWT)
	
	// Create services
	authService := service.NewAuthService(userRepo, authRepo, tenantRepo, jwtProvider, cfg.OAuth, cfg.Captcha)
	userService := service.NewUserService(userRepo, logger)
	
	// Create influencer service for the dashboard
	influencerService := service.NewInfluencerService(userRepo, referralRepo, logger, "https://ecomflex.com")
	
	// Create handlers
	authHandler := handler.NewAuthHandler(authService, logger)
	adminHandler := handler.NewAdminHandler(userService, logger)
	influencerHandler := handler.NewInfluencerHandler(influencerService, logger)
	
	// Create middleware
	authMiddleware := middleware.NewAuthMiddleware(authService, logger)
	tenantMiddleware := middleware.NewTenantMiddleware(logger)
	
	// API routes
	api := router.Group("/api")
	
	// Apply tenant middleware to all API routes
	api.Use(tenantMiddleware.ExtractTenantID())
	
	// Setup versioned routes (v1)
	v1 := api.Group("/v1")
	{
		// Authentication routes
		auth := v1.Group("/auth")
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.POST("/google/login", authHandler.GoogleLogin)
		auth.POST("/refresh", authHandler.RefreshToken)
		
		// Authenticated auth routes
		authProtected := auth.Group("/")
		authProtected.Use(authMiddleware.Authenticate())
		authProtected.POST("/logout", authHandler.Logout)
		
		// Influencer routes - this is our new addition
		influencer := v1.Group("/influencer")
		influencer.Use(authMiddleware.Authenticate())
		influencer.Use(authMiddleware.RequireRole("influencer"))
		
		// Dashboard data
		influencer.GET("/dashboard", influencerHandler.GetDashboardData)
		
		// Referral code
		influencer.GET("/referral-code", influencerHandler.GetReferralCode)
		
		// Referral statistics
		influencer.GET("/stats", influencerHandler.GetReferralStats)
		
		// Recent referrals
		influencer.GET("/recent-referrals", influencerHandler.GetRecentReferrals)
		
		// All referrals (with pagination)
		influencer.GET("/referrals", influencerHandler.GetRecentReferrals)
		
		// Admin routes
		admin := v1.Group("/admin")
		admin.Use(authMiddleware.Authenticate())
		admin.Use(authMiddleware.RequireRole("admin"))
		
		// Influencer management
		admin.GET("/influencers", adminHandler.GetInfluencers)
		admin.GET("/influencers/:id", adminHandler.GetInfluencer)
		admin.PUT("/influencers/:id/status", adminHandler.UpdateInfluencerStatus)
	}
	
	// Setup non-versioned routes for backward compatibility
	{
		// Authentication routes
		auth := api.Group("/auth")
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.POST("/google/login", authHandler.GoogleLogin)
		auth.POST("/refresh", authHandler.RefreshToken)
		
		// Authenticated auth routes
		authProtected := auth.Group("/")
		authProtected.Use(authMiddleware.Authenticate())
		authProtected.POST("/logout", authHandler.Logout)
		
		// Influencer routes - this is our new addition
		influencer := api.Group("/influencer")
		influencer.Use(authMiddleware.Authenticate())
		influencer.Use(authMiddleware.RequireRole("influencer"))
		
		// Dashboard data
		influencer.GET("/dashboard", influencerHandler.GetDashboardData)
		
		// Referral code
		influencer.GET("/referral-code", influencerHandler.GetReferralCode)
		
		// Referral statistics
		influencer.GET("/stats", influencerHandler.GetReferralStats)
		
		// Recent referrals
		influencer.GET("/recent-referrals", influencerHandler.GetRecentReferrals)
		
		// All referrals (with pagination)
		influencer.GET("/referrals", influencerHandler.GetRecentReferrals)
		
		// Admin routes
		admin := api.Group("/admin")
		admin.Use(authMiddleware.Authenticate())
		admin.Use(authMiddleware.RequireRole("admin"))
		
		// Influencer management
		admin.GET("/influencers", adminHandler.GetInfluencers)
		admin.GET("/influencers/:id", adminHandler.GetInfluencer)
		admin.PUT("/influencers/:id/status", adminHandler.UpdateInfluencerStatus)
	}
	
	return router
}