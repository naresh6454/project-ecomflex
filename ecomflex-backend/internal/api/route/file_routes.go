// internal/api/route/file_routes.go
package route

import (
    "github.com/gin-gonic/gin"
    "github.com/naresh6454/ecomflex-backend/internal/api/handler"
    "github.com/naresh6454/ecomflex-backend/internal/api/middleware"
)

func ConfigureFileRoutes(
    router *gin.RouterGroup,
    fileHandler *handler.FileHandler,
    authMiddleware *middleware.AuthMiddleware,
) {
    files := router.Group("/files")
    files.Use(authMiddleware.Authenticate()) // Ensure user is authenticated
    
    files.POST("/booking", fileHandler.UploadBookingDocument)
    files.POST("/review", fileHandler.UploadReviewMedia)
}