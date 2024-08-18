package usecases

import (
	"encoding/json"
	"os"

	"github.com/ertantorizkyf/gin-gorm-sample/dtos"
	"github.com/ertantorizkyf/gin-gorm-sample/helpers"
	"github.com/ertantorizkyf/gin-gorm-sample/models"
	"github.com/ertantorizkyf/gin-gorm-sample/repositories"
)

type ArticleUseCase struct {
	articleRepository repositories.ArticleRepository
}

func NewArticleUsecase(articleRepository repositories.ArticleRepository) ArticleUseCase {
	return ArticleUseCase{
		articleRepository: articleRepository,
	}
}

func (uc *ArticleUseCase) GetArticleList() []models.Article {
	articles := uc.articleRepository.GetArticleList()

	return articles
}

func (uc *ArticleUseCase) GetArticleById(id int) models.Article {
	article := uc.articleRepository.GetArticleByID(id)

	return article
}

func (uc *ArticleUseCase) CreateArticle(req dtos.ArticleReq) error {
	err := uc.articleRepository.CreateArticle(req)

	return err
}

func (uc *ArticleUseCase) UpdateArticle(id int, req dtos.ArticleReq) (models.Article, error) {
	article, err := uc.articleRepository.UpdateArticle(id, req)

	return article, err
}

func (uc *ArticleUseCase) DeleteArticle(id int) error {
	err := uc.articleRepository.DeleteArticle(id)

	return err
}

func (uc *ArticleUseCase) GetStructuredArticleJson() []dtos.Article {
	var articles []dtos.Article

	filePath := os.Getenv("ARTICLE_JSON_PATH")
	jsonRes := helpers.ReadJson(filePath)

	json.Unmarshal(jsonRes, &articles)

	return articles
}

func (uc *ArticleUseCase) GetUnstructuredArticleJson() []map[string]interface{} {
	var articles []map[string]interface{}

	filePath := os.Getenv("ARTICLE_JSON_PATH")
	jsonRes := helpers.ReadJson(filePath)

	json.Unmarshal([]byte(jsonRes), &articles)

	return articles
}
