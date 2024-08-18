package handlers

import (
	"strconv"

	"github.com/ertantorizkyf/gin-gorm-sample/dtos"
	"github.com/ertantorizkyf/gin-gorm-sample/models"
	"github.com/ertantorizkyf/gin-gorm-sample/usecases"
	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	articleUsecase usecases.ArticleUsecase
}

func NewArticleHandler(articleUsecase usecases.ArticleUsecase) ArticleHandler {
	return ArticleHandler{
		articleUsecase: articleUsecase,
	}
}

func (h *ArticleHandler) ArticleCreate(c *gin.Context) {
	body := dtos.ArticleReq{}
	c.Bind(&body)

	err := h.articleUsecase.CreateArticle(body)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})

		return
	}

	c.JSON(201, gin.H{
		"data": models.Article{Title: body.Title, Content: body.Content},
	})
}

func (h *ArticleHandler) ArticleIndex(c *gin.Context) {
	articles := h.articleUsecase.GetArticleList()

	c.JSON(200, gin.H{
		"data": articles,
	})
}

func (h *ArticleHandler) ArticleDetail(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	article := h.articleUsecase.GetArticleById(id)
	if article.ID == 0 {
		c.JSON(404, gin.H{
			"message": "ARTICLE_NOT_FOUND",
		})

		return
	}

	c.JSON(200, gin.H{
		"data": article,
	})
}

func (h *ArticleHandler) ArticleUpdate(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	body := dtos.ArticleReq{}
	c.Bind(&body)

	article, err := h.articleUsecase.UpdateArticle(id, body)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})

		return
	}

	c.JSON(200, gin.H{
		"data": article,
	})
}

func (h *ArticleHandler) ArticleDelete(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	err := h.articleUsecase.DeleteArticle(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})

		return
	}

	c.JSON(200, gin.H{
		"data": "Data deleted successfully",
	})
}
