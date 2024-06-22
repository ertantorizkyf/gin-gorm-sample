package helpers

import (
	"log"

	"github.com/redis/go-redis/v9"
)

func RedisClient() *redis.Client {
	url := "redis://127.0.0.1:6379"
	opts, err := redis.ParseURL(url)
	if err != nil {
		log.Fatal("[ERR] Failed to connect to Redis")
	} else {
		log.Printf("[INFO] Connected to Redis")
	}

	return redis.NewClient(opts)
}
