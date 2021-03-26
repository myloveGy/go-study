package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client
var ctx = context.Background()

func connection() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		Password: "",
		DB: 0,
	})
}

func Get(key string) (string, error) {
	return redisClient.Get(ctx, key).Result()
}

func Set(key, value string) error {
	return redisClient.Set(ctx, key, value, 0).Err()
}

func init() {
	connection()
}