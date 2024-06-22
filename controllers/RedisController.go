package controllers

import (
	"fmt"
	"log"

	"github.com/ertantorizkyf/gin-gorm-sample/helpers"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func RedisGet(c *gin.Context) {
	key := c.Query("key")

	rdb := helpers.RedisClient()

	redisVal, err := rdb.Get(c, key).Result()
	if err == redis.Nil {
		log.Printf("[ERR] Redis data of key \"%s\" not found", key)
		redisVal = ""
	} else if err != nil {
		log.Fatal("[ERR] Something went wrong while fetching redis key")
	}

	c.JSON(200, gin.H{
		"data": redisVal,
	})
}

func RedisSet(c *gin.Context) {
	var body struct {
		Key   string
		Value string
	}

	c.Bind(&body)

	rdb := helpers.RedisClient()

	err := rdb.Set(c, body.Key, body.Value, 0).Err()
	if err != nil {
		log.Fatal("[ERR] Something went wrong while setting redis key-value pair")
	}

	message := fmt.Sprintf("Successfully setting %s to %s", body.Value, body.Key)
	c.JSON(200, gin.H{
		"data": message,
	})
}
