package controllers

import (
	"encoding/json"
	"os"

	"github.com/ertantorizkyf/gin-gorm-sample/helpers"
	"github.com/ertantorizkyf/gin-gorm-sample/structs"
	"github.com/gin-gonic/gin"
)

func ArticleJsonStructuredIndex(c *gin.Context) {
	var articles []structs.Article

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
