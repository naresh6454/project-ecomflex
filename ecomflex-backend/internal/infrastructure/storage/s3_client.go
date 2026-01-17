package storage

import (
    "context"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "time"
    
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/credentials"
    "github.com/aws/aws-sdk-go-v2/service/s3"
    "github.com/aws/aws-sdk-go-v2/service/s3/types"
    "github.com/google/uuid"
    "github.com/naresh6454/ecomflex-backend/internal/infrastructure/config"
)

type S3Client struct {
    Client *s3.Client
    Config *config.S3Config
}

func NewS3Client(cfg *config.S3Config) *S3Client {
    // Log AWS configuration for debugging
    fmt.Printf("Initializing S3 client with: Region=%s, Bucket=%s\n", 
        cfg.Region, cfg.Bucket)
    
    // Create AWS configuration with credentials
    awsCfg := aws.Config{
        Region:      cfg.Region,
        Credentials: credentials.NewStaticCredentialsProvider(cfg.AccessKey, cfg.SecretKey, ""),
    }

    // Create S3 client
    client := s3.NewFromConfig(awsCfg)

    // Create the S3Client instance
    s3Client := &S3Client{
        Client: client,
        Config: cfg,
    }
    
    // Test AWS connection and permissions
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    
    _, err := client.ListBuckets(ctx, &s3.ListBucketsInput{})
    if err != nil {
        fmt.Printf("❌ AWS CONNECTION ERROR: %v\n", err)
        fmt.Printf("Please check your credentials and network connection\n")
    } else {
        fmt.Printf("✅ AWS connection successful\n")
        
        // Test bucket access
        _, err = client.HeadBucket(ctx, &s3.HeadBucketInput{
            Bucket: aws.String(cfg.Bucket),
        })
        if err != nil {
            fmt.Printf("❌ BUCKET ACCESS ERROR: %v\n", err)
            fmt.Printf("Please check if the bucket exists and you have permissions\n")
        } else {
            fmt.Printf("✅ Bucket '%s' is accessible\n", cfg.Bucket)
            
            // Configure bucket for public access
            err = s3Client.setupBucketForPublicAccess(ctx)
            if err != nil {
                fmt.Printf("⚠️ Warning: Failed to configure bucket for public access: %v\n", err)
                // Continue anyway as this might be due to IAM restrictions
            }
        }
    }

    return s3Client
}

// Set up the bucket for public access - configures both CORS and bucket policy
func (c *S3Client) setupBucketForPublicAccess(ctx context.Context) error {
    // Set up CORS
    corsErr := c.setupBucketCORS(ctx)
    if corsErr != nil {
        return fmt.Errorf("failed to set up CORS: %w", corsErr)
    }
    
    // Disable public access block
    blockErr := c.disablePublicAccessBlock(ctx)
    if blockErr != nil {
        return fmt.Errorf("failed to disable public access block: %w", blockErr)
    }
    
    // Set up bucket policy
    policyErr := c.setupBucketPolicy(ctx)
    if policyErr != nil {
        return fmt.Errorf("failed to set up bucket policy: %w", policyErr)
    }
    
    fmt.Printf("✅ Successfully configured bucket for public access\n")
    return nil
}

// Setup proper CORS configuration for the bucket
func (c *S3Client) setupBucketCORS(ctx context.Context) error {
    // Convert 86400 to int32 and create a pointer
    maxAgeSeconds := int32(86400) // 24 hours
    
    corsConfig := &s3.PutBucketCorsInput{
        Bucket: aws.String(c.Config.Bucket),
        CORSConfiguration: &types.CORSConfiguration{
            CORSRules: []types.CORSRule{
                {
                    AllowedHeaders: []string{"*"},
                    AllowedMethods: []string{"GET", "PUT", "POST", "DELETE", "HEAD"},
                    AllowedOrigins: []string{"*"}, // Allow all origins for development
                    ExposeHeaders:  []string{"ETag", "Content-Length", "Content-Type"},
                    MaxAgeSeconds:  aws.Int32(maxAgeSeconds), // Fixed: Use aws.Int32 for pointer conversion
                },
            },
        },
    }
    
    _, err := c.Client.PutBucketCors(ctx, corsConfig)
    if err != nil {
        fmt.Printf("❌ Failed to set CORS configuration: %v\n", err)
        return err
    }
    
    fmt.Printf("✅ CORS configuration set successfully\n")
    return nil
}

// Disable public access block settings
func (c *S3Client) disablePublicAccessBlock(ctx context.Context) error {
    // Attempt to remove public access block
    _, err := c.Client.DeletePublicAccessBlock(ctx, &s3.DeletePublicAccessBlockInput{
        Bucket: aws.String(c.Config.Bucket),
    })
    
    if err != nil {
        fmt.Printf("⚠️ Could not delete public access block: %v\n", err)
        // This may fail if the bucket doesn't have a public access block configuration
        // or if user doesn't have permissions, but we'll continue anyway
    } else {
        fmt.Printf("✅ Public access block removed\n")
    }
    
    // Wait a moment for the change to propagate
    time.Sleep(2 * time.Second)
    
    return nil
}

// Set up a bucket policy to allow public read access
func (c *S3Client) setupBucketPolicy(ctx context.Context) error {
    // Define a policy that allows public read access
    policyDocument := map[string]interface{}{
        "Version": "2012-10-17",
        "Statement": []map[string]interface{}{
            {
                "Sid":       "PublicReadGetObject",
                "Effect":    "Allow",
                "Principal": "*",
                "Action":    []string{"s3:GetObject"},
                "Resource":  fmt.Sprintf("arn:aws:s3:::%s/*", c.Config.Bucket),
            },
        },
    }
    
    // Convert policy to JSON
    policyJSON, err := json.Marshal(policyDocument)
    if err != nil {
        fmt.Printf("❌ Failed to create policy JSON: %v\n", err)
        return err
    }
    
    // Apply the policy to the bucket
    _, err = c.Client.PutBucketPolicy(ctx, &s3.PutBucketPolicyInput{
        Bucket: aws.String(c.Config.Bucket),
        Policy: aws.String(string(policyJSON)),
    })
    
    if err != nil {
        fmt.Printf("❌ Failed to set bucket policy: %v\n", err)
        return err
    }
    
    fmt.Printf("✅ Bucket policy set for public read access\n")
    return nil
}

// UploadFile uploads a file to S3 and returns its public URL
func (c *S3Client) UploadFile(ctx context.Context, fileData io.Reader, fileName string, contentType string) (string, error) {
    // Generate a unique filename to avoid collisions
    uniqueID := uuid.New().String()
    key := fmt.Sprintf("uploads/%s-%s", uniqueID, fileName)

    // Log upload details for debugging
    fmt.Printf("Attempting to upload to S3: bucket=%s, key=%s, contentType=%s, region=%s\n",
        c.Config.Bucket, key, contentType, c.Config.Region)

    // Upload the file to S3 with public-read ACL
    result, err := c.Client.PutObject(ctx, &s3.PutObjectInput{
        Bucket:      aws.String(c.Config.Bucket),
        Key:         aws.String(key),
        Body:        fileData,
        ContentType: aws.String(contentType),
        ACL:         types.ObjectCannedACLPublicRead,
        CacheControl: aws.String("max-age=31536000"), // Cache for 1 year
    })

    if err != nil {
        fmt.Printf("❌ S3 UPLOAD ERROR: %v\n", err)
        return "", fmt.Errorf("failed to upload file: %w", err)
    }

    fmt.Printf("✅ S3 upload successful: %+v\n", result)

    // Generate public URL based on region-specific format
    var fileURL string
    if c.Config.Region == "us-east-1" {
        // Special URL format for us-east-1 region
        fileURL = fmt.Sprintf("https://%s.s3.amazonaws.com/%s", 
            c.Config.Bucket, key)
    } else {
        // Standard URL format for all other regions
        fileURL = fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", 
            c.Config.Bucket, c.Config.Region, key)
    }

    fmt.Printf("Generated S3 URL: %s\n", fileURL)
    
    // Test if the URL is accessible
    req, _ := http.NewRequest("HEAD", fileURL, nil)
    client := &http.Client{Timeout: 5 * time.Second}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("⚠️ URL test failed: %v\n", err)
    } else {
        fmt.Printf("URL test result: %d %s\n", resp.StatusCode, resp.Status)
        resp.Body.Close()
        
        // If the URL is not accessible, try generating a pre-signed URL as fallback
        if resp.StatusCode == http.StatusForbidden || resp.StatusCode == http.StatusUnauthorized {
            fmt.Printf("⚠️ Public access failed, generating pre-signed URL as fallback\n")
            presignedURL, presignErr := c.GeneratePresignedURL(ctx, key, 24*time.Hour)
            if presignErr == nil {
                fmt.Printf("✅ Generated pre-signed URL: %s\n", presignedURL)
                return presignedURL, nil
            }
        }
    }
    
    return fileURL, nil
}

// GeneratePresignedURL generates a pre-signed URL for secure file access
func (c *S3Client) GeneratePresignedURL(ctx context.Context, key string, expires time.Duration) (string, error) {
    presignClient := s3.NewPresignClient(c.Client)
    
    presignResult, err := presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
        Bucket: aws.String(c.Config.Bucket),
        Key:    aws.String(key),
    }, s3.WithPresignExpires(expires))
    
    if err != nil {
        return "", fmt.Errorf("failed to generate presigned URL: %w", err)
    }
    
    return presignResult.URL, nil
}

// DeleteFile removes a file from S3
func (c *S3Client) DeleteFile(ctx context.Context, key string) error {
    _, err := c.Client.DeleteObject(ctx, &s3.DeleteObjectInput{
        Bucket: aws.String(c.Config.Bucket),
        Key:    aws.String(key),
    })
    
    if err != nil {
        return fmt.Errorf("failed to delete file: %w", err)
    }
    
    return nil
}