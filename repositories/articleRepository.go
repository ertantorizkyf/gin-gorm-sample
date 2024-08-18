package repositories

import (
	"github.com/ertantorizkyf/gin-gorm-sample/dtos"
	"github.com/ertantorizkyf/gin-gorm-sample/initializers"
	"github.com/ertantorizkyf/gin-gorm-sample/models"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	GetArticleList() []models.Article
	GetArticleByID(id int) models.Article
	CreateArticle(req dtos.ArticleReq) error
	UpdateArticle(id int, req dtos.ArticleReq) (models.Article, error)
	DeleteArticle(id int) error
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository() ArticleRepository {
	return &articleRepository{
		db: initializers.DB,
	}
}

func (r *articleRepository) GetArticleList() []models.Article {
	var articles []models.Article
	r.db.Find(&articles)

	return articles
}

func (r *articleRepository) GetArticleByID(id int) models.Article {
	var article models.Article
	r.db.First(&article, id)

	return article
}

func (r *articleRepository) CreateArticle(req dtos.ArticleReq) error {
	article := models.Article{Title: req.Title, Content: req.Content}

	trx := r.db.Begin()
	result := trx.Create(&article)
	if result.Error != nil {
		return result.Error
	}
	defer func() {
		trx.Rollback()
	}()

	result.Commit()

	return nil
}

func (r *articleRepository) UpdateArticle(id int, req dtos.ArticleReq) (models.Article, error) {
	var article models.Article
	r.db.First(&article, id)

	updatedArticle := models.Article{Title: req.Title, Content: req.Content}

	trx := r.db.Begin()
	result := trx.Model(&article).Updates(updatedArticle)
	if result.Error != nil {
		return article, result.Error
	}
	defer func() {
		trx.Rollback()
	}()

	result.Commit()

	return article, nil
}

func (r *articleRepository) DeleteArticle(id int) error {
	trx := r.db.Begin()
	result := trx.Delete(&models.Article{}, id)
	if result.Error != nil {
		return result.Error
	}
	defer func() {
		trx.Rollback()
	}()

	result.Commit()

	return nil
}
