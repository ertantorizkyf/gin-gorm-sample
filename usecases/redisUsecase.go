package usecases

import (
	_ "image/png"
	"log"

	"github.com/ertantorizkyf/gin-gorm-sample/dtos"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type RedisUsecase struct {
	rdb *redis.Client
}

func NewRedisUsecase(rdb *redis.Client) RedisUsecase {
	return RedisUsecase{
		rdb: rdb,
	}
}

func (uc *RedisUsecase) RedisGetKey(ctx *gin.Context, key string) (string, error) {
	redisVal, err := uc.rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		log.Printf("[ERR] Redis data of key \"%s\" not found", key)
		return redisVal, err
	} else if err != nil {
		log.Printf("[ERR] Something went wrong while fetching redis key")
		return redisVal, err
	}

	return redisVal, nil
}

func (uc *RedisUsecase) RedisSetKey(ctx *gin.Context, req dtos.RedisReq) error {
	err := uc.rdb.Set(ctx, req.Key, req.Value, 0).Err()
	if err != nil {
		log.Printf("[ERR] Something went wrong while setting redis key-value pair")
		return err
	}

	return nil
}
