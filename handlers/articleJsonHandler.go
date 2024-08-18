package handlers

import (
	"encoding/json"
	"os"

	"github.com/ertantorizkyf/gin-gorm-sample/dtos"
	"github.com/ertantorizkyf/gin-gorm-sample/helpers"
	"github.com/gin-gonic/gin"
)

func ArticleJsonStructuredIndex(c *gin.Context) {
	var articles []dtos.Article

	filePath := os.Getenv("ARTICLE_JSON_PATH")
	jsonRes := helpers.ReadJson(filePath)

	json.Unmarshal(jsonRes, &articles)

	c.JSON(200, gin.H{
		"data": articles,
	})
}

func ArticleJsonUnstructuredIndex(c *gin.Context) {
	var articles []map[string]interface{}

	filePath := os.Getenv("ARTICLE_JSON_PATH")
	jsonRes := helpers.ReadJson(filePath)

	json.Unmarshal([]byte(jsonRes), &articles)

	c.JSON(200, gin.H{
		"data": articles,
	})
}
