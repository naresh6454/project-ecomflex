package cache

import (
	"context"
	"fmt"

	redis "github.com/go-redis/redis/v8"
	"github.com/naresh6454/ecomflex-backend/config"
)

// RedisClient wraps redis.Client
type RedisClient struct {
	Client *redis.Client
}

// NewRedisClient creates a new Redis client
func NewRedisClient(cfg config.RedisConfig) (*RedisClient, error) {
	// Create client
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	
	// Verify connection
	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}
	
	return &RedisClient{
		Client: client,
	}, nil
}

// Close closes the Redis client
func (c *RedisClient) Close() error {
	return c.Client.Close()
}