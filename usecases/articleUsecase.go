package usecases

import (
	"encoding/json"
	"os"

	"github.com/ertantorizkyf/gin-gorm-sample/dtos"
	"github.com/ertantorizkyf/gin-gorm-sample/helpers"
	"github.com/ertantorizkyf/gin-gorm-sample/models"
	"github.com/ertantorizkyf/gin-gorm-sample/repositories"
)

type ArticleUsecase struct {
	articleRepository repositories.ArticleRepository
}

func NewArticleUsecase(articleRepository repositories.ArticleRepository) ArticleUsecase {
	return ArticleUsecase{
		articleRepository: articleRepository,
	}
}

func (uc *ArticleUsecase) GetArticleList() []models.Article {
	articles := uc.articleRepository.GetArticleList()

	return articles
}

func (uc *ArticleUsecase) GetArticleById(id int) models.Article {
	article := uc.articleRepository.GetArticleByID(id)

	return article
}

func (uc *ArticleUsecase) CreateArticle(req dtos.ArticleReq) error {
	err := uc.articleRepository.CreateArticle(req)

	return err
}

func (uc *ArticleUsecase) UpdateArticle(id int, req dtos.ArticleReq) (models.Article, error) {
	article, err := uc.articleRepository.UpdateArticle(id, req)

	return article, err
}

func (uc *ArticleUsecase) DeleteArticle(id int) error {
	err := uc.articleRepository.DeleteArticle(id)

	return err
}

func (uc *ArticleUsecase) GetStructuredArticleJson() []dtos.Article {
	var articles []dtos.Article

	filePath := os.Getenv("ARTICLE_JSON_PATH")
	jsonRes := helpers.ReadJson(filePath)

	json.Unmarshal(jsonRes, &articles)

	return articles
}

func (uc *ArticleUsecase) GetUnstructuredArticleJson() []map[string]interface{} {
	var articles []map[string]interface{}

	filePath := os.Getenv("ARTICLE_JSON_PATH")
	jsonRes := helpers.ReadJson(filePath)

	json.Unmarshal([]byte(jsonRes), &articles)

	return articles
}
