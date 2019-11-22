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

func (category *CategoriesModel) TableName() string {
	return "categories"
}

// 创建新标签
func (category *CategoriesModel) Create() error {
	return DB.Self.Create(&category).Error
}

// 获取全部标签
func (category *CategoriesModel) ListCategories(offset, limit int) ([]*CategoriesModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	categories := make([]*CategoriesModel, 0)
	var count uint64

	if err := DB.Self.Model(&category).Count(&count).Error; err != nil {
		return categories, count, err
	}
	if err := DB.Self.Where("").Offset(offset).Limit(limit).Order("id desc").Find(&categories).Error; err != nil {
		return categories, count, err
	}

	return categories, count, nil
}

// 根据标签id获取标签数据.
func (category *CategoriesModel) GetCategoryById(id uint64) (*CategoriesModel, error) {
	d := DB.Self.Where("id = ?", id).First(&category)
	return category, d.Error
}

// 根据标签id删除标签
func (category *CategoriesModel)  DeleteCategory(id uint64) error {
	category.BaseModel.Id = id
	return DB.Self.Delete(&category).Error
}