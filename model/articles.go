package model

import (
	"ant-forum/pkg/constvar"
	"gopkg.in/go-playground/validator.v9"
	"time"
)

type ArticleModel struct {
	BaseModel
	Title      string `json:"title" gorm:"column:title;not null" binding:"required"`
	Content    string `json:"content" gorm:"column:content;not null" binding:"required"`
	CategoryId uint64 `json:"category_id" gorm:"column:category_id;not null" binding:"required"`
	TagId      uint64 `json:"tag_id" gorm:"column:tag_id;not null" binding:"required"`
	UserId     uint64 `json:"user_id" gorm:"column:user_id;not null" binding:"required"`
}

type ArticleInfo struct {
	Id           uint64    `json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	CategoryId   uint64    `json:"category_id"`
	CategoryName string    `json:"category_name"`
	TagId        uint64    `json:"tag_id"`
	TagName      string    `json:"tag_name"`
	UserId       uint64    `json:"user_id"`
	UserName     string    `json:"username"`
	Avatar       string    `json:"avatar"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"update_at"`
}

func (c *ArticleModel) TableName() string {
	return "articles"
}

// 创建新标签
func (t *ArticleModel) Create() error {
	return DB.Self.Create(&t).Error
}

// 根据标签id获取文章.
func GetArticleById(id uint64) (*ArticleModel, error) {
	u := &ArticleModel{}
	d := DB.Self.First(&u, id)
	return u, d.Error
}

// 根据标签id删除文章
func DeleteArticle(id uint64) error {
	article := ArticleModel{}
	article.BaseModel.Id = id
	return DB.Self.Delete(&article).Error
}

// 获取文章分页
func ListArticles(offset, limit int) ([]*ArticleModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}
	articles := make([]*ArticleModel, 0)
	var count uint64

	if err := DB.Self.Model(&ArticleModel{}).Count(&count).Error; err != nil {
		return articles, count, err
	}
	if err := DB.Self.Where("").Offset(offset).Limit(limit).Order("id desc").Find(&articles).Error; err != nil {
		return articles, count, err
	}

	return articles, count, nil
}

// 验证创建字段
func (t *ArticleModel) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}
