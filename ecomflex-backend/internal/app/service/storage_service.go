// internal/app/service/storage_service.go
package service

import (
    "context"
    "fmt"
    "mime/multipart"
    "path/filepath"
    "strings"
    "time"
    
    "github.com/google/uuid"
    "github.com/naresh6454/ecomflex-backend/internal/infrastructure/storage"
)

type StorageService interface {
    UploadFile(ctx context.Context, file *multipart.FileHeader, directory string) (string, error)
    GetFileURL(ctx context.Context, key string, expires time.Duration) (string, error)
    DeleteFile(ctx context.Context, key string) error
}

type storageServiceImpl struct {
    s3Client *storage.S3Client
}

func NewStorageService(s3Client *storage.S3Client) StorageService {
    return &storageServiceImpl{
        s3Client: s3Client,
    }
}

// UploadFile handles file upload from multipart form
func (s *storageServiceImpl) UploadFile(ctx context.Context, file *multipart.FileHeader, directory string) (string, error) {
    // Open the uploaded file
    src, err := file.Open()
    if err != nil {
        return "", fmt.Errorf("failed to open file: %w", err)
    }
    defer src.Close()
    
    // Get content type
    contentType := file.Header.Get("Content-Type")
    if contentType == "" {
        // Determine content type from extension
        ext := strings.ToLower(filepath.Ext(file.Filename))
        switch ext {
        case ".jpg", ".jpeg":
            contentType = "image/jpeg"
        case ".png":
            contentType = "image/png"
        case ".pdf":
            contentType = "application/pdf"
        case ".mp4":
            contentType = "video/mp4"
        default:
            contentType = "application/octet-stream"
        }
    }
    
    // Create a key with directory prefix
    fileName := file.Filename
    uniqueID := uuid.New().String()
    key := fmt.Sprintf("%s/%s-%s", directory, uniqueID, fileName)
    
    // Upload to S3
    fileURL, err := s.s3Client.UploadFile(ctx, src, key, contentType)
    if err != nil {
        return "", err
    }
    
    return fileURL, nil
}

// GetFileURL returns a pre-signed URL for secure access
func (s *storageServiceImpl) GetFileURL(ctx context.Context, key string, expires time.Duration) (string, error) {
    return s.s3Client.GeneratePresignedURL(ctx, key, expires)
}

// DeleteFile removes a file
func (s *storageServiceImpl) DeleteFile(ctx context.Context, key string) error {
    return s.s3Client.DeleteFile(ctx, key)
}