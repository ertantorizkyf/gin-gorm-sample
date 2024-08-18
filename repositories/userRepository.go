package repositories

import (
	"github.com/ertantorizkyf/gin-gorm-sample/dtos"
	"github.com/ertantorizkyf/gin-gorm-sample/initializers"
	"github.com/ertantorizkyf/gin-gorm-sample/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetArticleList() []models.Article
	GetUserByEmail(email string) models.User
	RegisterUser(req dtos.RegisterReq) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{
		db: initializers.DB,
	}
}

func (r *userRepository) GetArticleList() []models.Article {
	var articles []models.Article
	r.db.Find(&articles)

	return articles
}

func (r *userRepository) GetUserByEmail(email string) models.User {
	var user models.User
	r.db.Where("email = ?", email).First(&user)

	return user
}

func (r *userRepository) RegisterUser(req dtos.RegisterReq) error {
	user := models.User{FullName: req.FullName, Email: req.Email, Password: req.Password}

	trx := r.db.Begin()
	result := trx.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	defer func() {
		trx.Rollback()
	}()

	result.Commit()

	return nil
}
