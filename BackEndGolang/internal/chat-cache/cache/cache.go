package cache

import "github.com/go-redis/redis/v8"

type Cache struct {
	Client *redis.Client
	SecretKey   string
}

func NewRedisClient(redisAddr string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})
}
