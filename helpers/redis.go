package helpers

import (
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

func RedisClient() *redis.Client {
	url := os.Getenv("REDIS_URL")
	opts, err := redis.ParseURL(url)
	if err != nil {
		log.Fatal("[ERR] Failed to connect to Redis")
	} else {
		log.Printf("[INFO] Connected to Redis")
	}

	return redis.NewClient(opts)
}
