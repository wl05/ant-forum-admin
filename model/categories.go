package model

import (
	"ant-forum/pkg/constvar"
)

type CategoriesModel struct {
	BaseModel
	CategoryName string `json:"category_name" gorm:"column:category_name;not null" binding:"required"`
}

type CategoryInfo struct {
	Id           uint64 `json:"id"`
	CategoryName string `json:"category_name"`
}

func (c *CategoriesModel) TableName() string {
	return "categories"
}

// 创建新标签
func (c *CategoriesModel) Create() error {
	return DB.Self.Create(&c).Error
}

// 获取全部标签
func (c *CategoriesModel) ListCategories(offset, limit int) ([]*CategoriesModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	categories := make([]*CategoriesModel, 0)
	var count uint64

	if err := DB.Self.Model(&c).Count(&count).Error; err != nil {
		return categories, count, err
	}
	if err := DB.Self.Where("").Offset(offset).Limit(limit).Order("id desc").Find(&categories).Error; err != nil {
		return categories, count, err
	}

	return categories, count, nil
}

// 根据标签id获取标签数据.
func (c *CategoriesModel) GetCategoryById(id uint64) (*CategoriesModel, error) {
	d := DB.Self.Where("id = ?", id).First(&c)
	return c, d.Error
}

// 根据标签id删除标签
func (c *CategoriesModel) DeleteCategory(id uint64) error {
	c.BaseModel.Id = id
	return DB.Self.Delete(&c).Error
}
