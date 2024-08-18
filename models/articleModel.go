package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title   string
	Content string
}

func (Article) TableName() string {
	return "articles"
}
