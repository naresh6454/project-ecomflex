package handler

import (
    "fmt"
    "net/http"
    "path/filepath"
    "strconv"
    "time"
    
    "github.com/gin-gonic/gin"
    "github.com/naresh6454/ecomflex-backend/internal/api/response"
    "github.com/naresh6454/ecomflex-backend/internal/app/service"
)

type FileHandler struct {
    storageService service.StorageService
}

func NewFileHandler(storageService service.StorageService) *FileHandler {
    return &FileHandler{
        storageService: storageService,
    }
}

// UploadBookingDocument handles booking document uploads
func (h *FileHandler) UploadBookingDocument(c *gin.Context) {
    // Get user ID from context (set by middleware)
    userID, exists := c.Get("userID")
    if !exists {
        response.Error(c, http.StatusUnauthorized, "Unauthorized", nil)
        return
    }
    
    // Get product ID
    productID := c.PostForm("productId")
    if productID == "" {
        response.Error(c, http.StatusBadRequest, "Product ID is required", nil)
        return
    }
    
    // Get the uploaded file
    file, err := c.FormFile("file")
    if err != nil {
        response.Error(c, http.StatusBadRequest, "Invalid file", nil)
        return
    }
    
    // Log file info for debugging
    fmt.Printf("Uploading file: name=%s, size=%d, content-type=%s\n", 
        file.Filename, file.Size, file.Header.Get("Content-Type"))
    
    // Directory for booking documents
    directory := fmt.Sprintf("bookings/%s/%s", userID.(string), productID)
    
    // Upload file
    fileURL, err := h.storageService.UploadFile(c, file, directory)
    if err != nil {
        fmt.Printf("Failed to upload file: %v\n", err)
        response.Error(c, http.StatusInternalServerError, "Failed to upload file", err)
        return
    }
    
    fmt.Printf("File uploaded successfully: %s\n", fileURL)
    
    // Add CORS headers for direct response handling
    c.Header("Access-Control-Allow-Origin", "*")
    c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
    c.Header("Access-Control-Allow-Headers", "Content-Type")
    
    // Return success response
    response.Success(c, http.StatusOK, "File uploaded successfully", gin.H{
        "fileUrl": fileURL,
        "fileName": file.Filename,
        "fileSize": file.Size,
        "fileType": file.Header.Get("Content-Type"),
        "s3Path": directory + "/" + filepath.Base(fileURL),
    })
}

// UploadReviewMedia handles review media uploads
func (h *FileHandler) UploadReviewMedia(c *gin.Context) {
    // Get user ID from context
    userID, exists := c.Get("userID")
    if !exists {
        response.Error(c, http.StatusUnauthorized, "Unauthorized", nil)
        return
    }
    
    // Get product ID
    productID := c.PostForm("productId")
    if productID == "" {
        response.Error(c, http.StatusBadRequest, "Product ID is required", nil)
        return
    }
    
    // Get review text and rating
    reviewText := c.PostForm("text")
    ratingStr := c.PostForm("rating")
    
    // Get the uploaded file(s)
    form, err := c.MultipartForm()
    if err != nil {
        response.Error(c, http.StatusBadRequest, "Invalid form data", nil)
        return
    }
    
    files := form.File["files[]"]
    if len(files) == 0 && reviewText == "" {
        response.Error(c, http.StatusBadRequest, "No files or review text provided", nil)
        return
    }
    
    fmt.Printf("Uploading %d review media files for product %s\n", len(files), productID)
    
    // Directory for review media
    directory := fmt.Sprintf("reviews/%s/%s", userID.(string), productID)
    
    fileURLs := make([]string, 0, len(files))
    fileDetails := make([]map[string]interface{}, 0, len(files))
    
    // Upload each file
    for _, file := range files {
        fileURL, err := h.storageService.UploadFile(c, file, directory)
        if err != nil {
            fmt.Printf("Failed to upload review file: %v\n", err)
            response.Error(c, http.StatusInternalServerError, "Failed to upload file", err)
            return
        }
        
        fileURLs = append(fileURLs, fileURL)
        fileDetails = append(fileDetails, map[string]interface{}{
            "url": fileURL,
            "name": file.Filename,
            "size": file.Size,
            "type": file.Header.Get("Content-Type"),
        })
    }
    
    // Add CORS headers for direct response handling
    c.Header("Access-Control-Allow-Origin", "*")
    c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
    c.Header("Access-Control-Allow-Headers", "Content-Type")
    
    // Return success response with additional debugging info
    response.Success(c, http.StatusOK, "Files uploaded successfully", gin.H{
        "fileUrls": fileURLs,
        "fileDetails": fileDetails,
        "text": reviewText,
        "rating": ratingStr,
        "uploadDirectory": directory,
    })
}

// GetPresignedURL generates a temporary URL for accessing files
func (h *FileHandler) GetPresignedURL(c *gin.Context) {
    // Get the key parameter
    key := c.Query("key")
    if key == "" {
        response.Error(c, http.StatusBadRequest, "Key parameter is required", nil)
        return
    }
    
    // Set expiration time - 24 hours by default
    var expiration time.Duration = 24 * time.Hour
    
    // Check if expiration time is specified
    expirationStr := c.DefaultQuery("expiration", "24")
    if expirationInt, err := strconv.Atoi(expirationStr); err == nil && expirationInt > 0 {
        expiration = time.Duration(expirationInt) * time.Hour
    }
    
    // Generate presigned URL
    presignedURL, err := h.storageService.GetFileURL(c.Request.Context(), key, expiration)
    if err != nil {
        response.Error(c, http.StatusInternalServerError, "Failed to generate presigned URL", err)
        return
    }
    
    // Return success response
    response.Success(c, http.StatusOK, "Presigned URL generated successfully", gin.H{
        "url": presignedURL,
        "expiration": expirationStr + " hours",
        "expiresAt": time.Now().Add(expiration).Format(time.RFC3339),
    })
}