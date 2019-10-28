package model

import (
	"gopkg.in/go-playground/validator.v9"
)

type ArticleModel struct {
	BaseModel
	Title      string `json:"title" gorm:"column:title;not null" binding:"required"`
	Content    string `json:"content" gorm:"column:content;not null" binding:"required"`
	CategoryId uint64 `json:"category_id" gorm:"column:category_id;not null" binding:"required"`
	TagId      uint64 `json:"tag_id" gorm:"column:tag_id;not null" binding:"required"`
	UserId     uint64 `json:"user_id" gorm:"column:user_id;not null" binding:"required"`
}

func (c *ArticleModel) TableName() string {
	return "articles"
}

// 创建新标签
func (t *ArticleModel) Create() error {
	return DB.Self.Create(&t).Error
}

// 验证创建字段
func (t *ArticleModel) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}
