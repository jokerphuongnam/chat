package cache

import (
	"chat-service/internal/logs"
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

func (cache *Cache) CacheJWTToken(userID string, jwtToken string) error {
	// Set expiration duration for JWT token
	tokenExpiration := time.Hour * 24

	// Generate unique ID for this token
	tokenID := uuid.New().String()
	redisKey := "auth:jwt:" + userID + ":" + tokenID

	// Cache the JWT token with an expiration time in Redis
	err := cache.Client.Set(context.Background(), redisKey, jwtToken, tokenExpiration).Err()
	if err != nil {
		return fmt.Errorf("failed to cache JWT token: %v", err)
	}

	logs.Log.Printf("key: %v -- token: %v\n", redisKey, jwtToken)

	// Add this token to the user's sorted set for cleanup purposes
	expirationTimestamp := time.Now().Add(tokenExpiration).Unix()
	cache.Client.ZAdd(context.Background(), "auth:tokens:"+userID, &redis.Z{
		Score:  float64(expirationTimestamp),
		Member: redisKey,
	})

	return nil
}
