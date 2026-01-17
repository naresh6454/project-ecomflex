package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/naresh6454/ecomflex-backend/config"
)

// CorsMiddleware configures CORS for the application
func CorsMiddleware(cfg *config.CorsConfig) gin.HandlerFunc {
	// Make sure the X-Superadmin-Auth header is included
	allowHeaders := append(cfg.AllowHeaders, "X-Superadmin-Auth")
	
	// Ensure no duplicates in the headers list
	headerMap := make(map[string]bool)
	var uniqueHeaders []string
	
	for _, header := range allowHeaders {
		if !headerMap[header] {
			headerMap[header] = true
			uniqueHeaders = append(uniqueHeaders, header)
		}
	}
	
	// Configure CORS
	corsConfig := cors.Config{
		AllowOrigins:     cfg.AllowOrigins,
		AllowMethods:     cfg.AllowMethods,
		AllowHeaders:     uniqueHeaders, // Use the enhanced headers list
		ExposeHeaders:    cfg.ExposeHeaders,
		AllowCredentials: cfg.AllowCredentials,
		MaxAge:           cfg.MaxAge,
	}
	
	return cors.New(corsConfig)
}