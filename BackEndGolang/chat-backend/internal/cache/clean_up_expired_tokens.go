package cache

import (
	"context"
	"fmt"
	"time"
)

func (cache *Cache) CleanUpExpiredTokens(userID string) error {
	currentTime := float64(time.Now().Unix())

	cache.Client.ZRemRangeByScore(context.Background(), "auth:tokens:"+userID, "-inf", fmt.Sprintf("%f", currentTime))

	pattern := "auth:jwt:" + userID + ":*"

	iter := cache.Client.Scan(context.Background(), 0, pattern, 0).Iterator()

	for iter.Next(context.Background()) {
		key := iter.Val()

		_, err := cache.Client.TTL(context.Background(), key).Result()

		if err != nil {
			return err
		}

		err = cache.Client.Del(context.Background(), key).Err()
		if err != nil {
			return err
		}
	}

	if err := iter.Err(); err != nil {
		return err
	}

	return nil
}
