package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/ertantorizkyf/gin-gorm-sample/structs"
	"github.com/gin-gonic/gin"
)

func ArticleJsonStructuredIndex(c *gin.Context) {
	var articles []structs.Article

	filePath := os.Getenv("ARTICLE_JSON_PATH")
	jsonFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal("[ERR] Failed to read json file")
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &articles)

	c.JSON(200, gin.H{
		"data": articles,
	})
}

func ArticleJsonUnstructuredIndex(c *gin.Context) {
	var articles []map[string]interface{}

	filePath := os.Getenv("ARTICLE_JSON_PATH")
	jsonFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal("[ERR] Failed to read json file")
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &articles)

	c.JSON(200, gin.H{
		"data": articles,
	})
}
