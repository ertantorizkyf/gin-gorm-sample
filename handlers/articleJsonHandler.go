package handlers

import (
	"github.com/ertantorizkyf/gin-gorm-sample/usecases"
	"github.com/gin-gonic/gin"
)

type ArticleJsonHandler struct {
	articleUsecase usecases.ArticleUsecase
}

func NewArticleJsonHandler(articleUsecase usecases.ArticleUsecase) ArticleJsonHandler {
	return ArticleJsonHandler{
		articleUsecase: articleUsecase,
	}
}

func (h *ArticleJsonHandler) ArticleJsonStructuredIndex(c *gin.Context) {
	articles := h.articleUsecase.GetStructuredArticleJson()

	c.JSON(200, gin.H{
		"data": articles,
	})
}

func (h *ArticleJsonHandler) ArticleJsonUnstructuredIndex(c *gin.Context) {
	articles := h.articleUsecase.GetUnstructuredArticleJson()

	c.JSON(200, gin.H{
		"data": articles,
	})
}
