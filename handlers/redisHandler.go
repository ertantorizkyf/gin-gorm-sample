package handlers

import (
	"fmt"

	"github.com/ertantorizkyf/gin-gorm-sample/dtos"
	"github.com/ertantorizkyf/gin-gorm-sample/usecases"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type RedisHandler struct {
	redisUsecase usecases.RedisUsecase
}

func NewRedisHandler(redisUsecase usecases.RedisUsecase) RedisHandler {
	return RedisHandler{
		redisUsecase: redisUsecase,
	}
}

func (h *RedisHandler) RedisGet(c *gin.Context) {
	key := c.Query("key")

	redisVal, err := h.redisUsecase.RedisGetKey(c, key)
	if err != nil {
		statusCode := 500
		if err == redis.Nil {
			statusCode = 404
		}

		c.JSON(statusCode, gin.H{
			"message": err,
		})

		return
	}

	c.JSON(200, gin.H{
		"data": redisVal,
	})
}

func (h *RedisHandler) RedisSet(c *gin.Context) {
	body := dtos.RedisReq{}
	c.Bind(&body)

	err := h.redisUsecase.RedisSetKey(c, body)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})

		return
	}

	message := fmt.Sprintf("Successfully setting %s to %s", body.Value, body.Key)
	c.JSON(200, gin.H{
		"data": message,
	})
}
