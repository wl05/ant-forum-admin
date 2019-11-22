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

func (tag *TagModel) TableName() string {
	return "tags"
}

// 创建新标签
func (tag *TagModel) Create() error {
	return DB.Self.Create(&tag).Error
}

// 获取全部标签
func (tag *TagModel) ListTags(offset, limit int) ([]*TagModel, uint64, error) {
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
func (tag *TagModel) GetTagById(id uint64) (*TagModel, error) {
	d := DB.Self.First(&tag, id)
	return tag, d.Error
}

// 根据标签id删除标签
func (tag *TagModel) DeleteTag(id uint64) error {
	tag.BaseModel.Id = id
	return DB.Self.Delete(&tag).Error
}

// 更新标签
func (tag *TagModel) Update() error {
	return DB.Self.Omit("created_at").Save(tag).Error
}
