package model

import (
	"ant-forum/pkg/constvar"
)

type TagModel struct {
	BaseModel
	TagName string `form:"tag_name" json:"tag_name" gorm:"column:tag_name;not null" binding:"required"`
}

type TagInfo struct {
	Id      uint64 `json:"id"`
	TagName string `json:"tag_name"`
}

func (t *TagModel) TableName() string {
	return "tags"
}

// 创建新标签
func (t *TagModel) Create() error {
	return DB.Self.Create(&t).Error
}

// 获取全部标签
func (t *TagModel) ListTags(offset, limit int) ([]*TagModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	tags := make([]*TagModel, 0)
	var count uint64
	if err := DB.Self.Model(&TagModel{}).Count(&count).Error; err != nil {
		return tags, count, err
	}
	if err := DB.Self.Where("").Offset(offset).Limit(limit).Order("id desc").Find(&tags).Error; err != nil {
		return tags, count, err
	}

	return tags, count, nil
}

// 根据标签id获取标签数据.
func (t *TagModel) GetTagById(id uint64) (*TagModel, error) {
	d := DB.Self.First(&t, id)
	return t, d.Error
}

// 根据标签id删除标签
func (t *TagModel) DeleteTag(id uint64) error {
	t.BaseModel.Id = id
	return DB.Self.Delete(&t).Error
}

// 更新标签
func (t *TagModel) Update() error {
	return DB.Self.Omit("created_at").Save(t).Error
}
