package controllers

import (
	"github.com/ertantorizkyf/gin-gorm-sample/initializers"
	"github.com/ertantorizkyf/gin-gorm-sample/models"
	"github.com/gin-gonic/gin"
)

func ArticleCreate(c *gin.Context) {
	var body struct {
		Title   string
		Content string
	}

	c.Bind(&body)

	article := models.Article{Title: body.Title, Content: body.Content}
	result := initializers.DB.Create(&article)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"data": article,
	})
}

func ArticleIndex(c *gin.Context) {
	var articles []models.Article
	initializers.DB.Find(&articles)

	c.JSON(200, gin.H{
		"data": articles,
	})
}

func ArticleDetail(c *gin.Context) {
	id := c.Param("id")

	var article models.Article
	initializers.DB.First(&article, id)

	c.JSON(200, gin.H{
		"data": article,
	})
}

func ArticleUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Title   string
		Content string
	}

	c.Bind(&body)

	var article models.Article
	initializers.DB.First(&article, id)

	updatedArticle := models.Article{Title: body.Title, Content: body.Content}
	initializers.DB.Model(&article).Updates(updatedArticle)

	c.JSON(200, gin.H{
		"data": article,
	})
}

func ArticleDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Article{}, id)

	c.JSON(200, gin.H{
		"data": "Data deleted successfully",
	})
}
