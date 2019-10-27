package model

import (
	"ant-forum/pkg/constvar"
	validator "gopkg.in/go-playground/validator.v9"
	"sync"
)

type TagModel struct {
	BaseModel
	TagName string `json:"tag_name" gorm:"column:tag_name;not null" binding:"required"`
}

type TagInfo struct {
	Id      uint64 `json:"id"`
	TagName string `json:"tag_name"`
}

type TagsList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*TagInfo
}

func (c *TagModel) TableName() string {
	return "tags"
}

// 创建新标签
func (t *TagModel) Create() error {
	return DB.Self.Create(&t).Error
}

// 获取全部标签
func ListTags(offset, limit int) ([]*TagModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	tags := make([]*TagModel, 0)
	var count uint64
	if err := DB.Self.Where("").Offset(offset).Limit(limit).Order("id desc").Find(&tags).Error; err != nil {
		return tags, count, err
	}

	return tags, count, nil
}

// 根据标签id获取标签数据.
func GetTagById(id uint64) (*TagModel, error) {
	u := &TagModel{}
	d := DB.Self.First(&u, id)
	return u, d.Error
}

// 验证创建字段
func (t *TagModel) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}
