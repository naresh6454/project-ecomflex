package config

import (
    "log"
    "os"
)

type S3Config struct {
    AccessKey string
    SecretKey string
    Region    string
    Bucket    string
}

func NewS3Config() *S3Config {
    accessKey := os.Getenv("AWS_ACCESS_KEY")
    secretKey := os.Getenv("AWS_SECRET_KEY")
    region := os.Getenv("AWS_REGION")
    bucket := os.Getenv("AWS_S3_BUCKET")

    if accessKey == "" || secretKey == "" || region == "" || bucket == "" {
        log.Fatal("Missing one or more required AWS environment variables")
    }

    log.Printf("✅ S3 Config Loaded | Region: %s | Bucket: %s\n", region, bucket)

    return &S3Config{
        AccessKey: accessKey,
        SecretKey: secretKey,
        Region:    region,
        Bucket:    bucket,
    }
}
