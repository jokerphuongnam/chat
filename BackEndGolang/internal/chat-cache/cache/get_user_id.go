package cache

import (
	"chat-logs/logs"
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

func (cache *Cache) GetUserID(jwtToken string) (uuid.UUID, error) {
	ctx := context.Background()

	// Use Redis SCAN to search for the key that has jwtToken as its value.
	iter := cache.Client.Scan(ctx, 0, "auth:jwt:*", 0).Iterator()
	for iter.Next(ctx) {
		redisKey := iter.Val()

		// Retrieve the value stored at this key
		storedToken, err := cache.Client.Get(ctx, redisKey).Result()
		if err != nil {
			if err == redis.Nil {
				continue // Token not found at this key, keep scanning
			}
			return uuid.Nil, fmt.Errorf("redis error: %v", err)
		}

		// Check if the stored token matches the jwtToken we're looking for
		if storedToken == jwtToken {
			// Extract userID from the key format: "auth:jwt:<userID>:<tokenID>"
			parts := strings.Split(redisKey, ":")
			if len(parts) < 3 {
				return uuid.Nil, fmt.Errorf("invalid key format: %s", redisKey)
			}

			userIDStr := parts[2]
			logs.Log.Infof("userID: %s - redisKey: %s", userIDStr, redisKey)
			userID, err := uuid.Parse(userIDStr)
			if err != nil {
				return uuid.Nil, fmt.Errorf("invalid user ID format in key: %s", err)
			}

			return userID, nil
		}
	}

	if err := iter.Err(); err != nil {
		return uuid.Nil, fmt.Errorf("error during Redis scan: %v", err)
	}

	// If no matching token was found
	return uuid.Nil, errors.New("token not found in Redis cache")
}
