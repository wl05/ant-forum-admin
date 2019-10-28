package model

import (
	"ant-forum/pkg/constvar"
	validator "gopkg.in/go-playground/validator.v9"
	"sync"
)

type CategoriesModel struct {
	BaseModel
	CategoryName string `json:"category_name" gorm:"column:category_name;not null" binding:"required"`
}

type CategoryInfo struct {
	Id           uint64 `json:"id"`
	CategoryName string `json:"category_name"`
}

type CategoriesList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*CategoryInfo
}

func (c *CategoriesModel) TableName() string {
	return "categories"
}

// 创建新标签
func (t *CategoriesModel) Create() error {
	return DB.Self.Create(&t).Error
}

// 获取全部标签
func ListCategories(offset, limit int) ([]*CategoriesModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	categories := make([]*CategoriesModel, 0)
	var count uint64
	if err := DB.Self.Where("").Offset(offset).Limit(limit).Order("id desc").Find(&categories).Error; err != nil {
		return categories, count, err
	}

	return categories, count, nil
}

// 根据标签id获取标签数据.
func GetCategoryById(id uint64) (*CategoriesModel, error) {
	u := &CategoriesModel{}
	d := DB.Self.First(&u, id)
	return u, d.Error
}

// 根据标签id删除标签
func DeleteCategory(id uint64) error {
	category := CategoriesModel{}
	category.BaseModel.Id = id
	return DB.Self.Delete(&category).Error
}

// 验证创建字段
func (t *CategoriesModel) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}
