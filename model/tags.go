package model

type TagModel struct {
	BaseModel
	TagName string `json:"tag_name" gorm:"column:tag_name;not null" binding:"required"`
}


func (c *TagModel) TableName() string {
	return "tags"
}