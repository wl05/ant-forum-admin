package model

import (
	validator "gopkg.in/go-playground/validator.v9"
)

type TagModel struct {
	BaseModel
	TagName string `json:"tag_name" gorm:"column:tag_name;not null" binding:"required"`
}

func (c *TagModel) TableName() string {
	return "tags"
}

// Create creates a new tag.
func (t *TagModel) Create() error {
	return DB.Self.Create(&t).Error
}

// Validate the fields.
func (t *TagModel) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}
