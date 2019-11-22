package model

import (
	"ant-forum/pkg/auth"
	"ant-forum/pkg/constvar"
	"fmt"
)

// User represents a registered user.
type UserModel struct {
	BaseModel
	Avatar   string `json:"avatar" gorm:"column:avatar;not null" binding:"required"`
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}
type UserInfo struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

func (user *UserModel) TableName() string {
	return "users"
}

// Create creates a new user account.
func (user *UserModel) Create() error {
	return DB.Self.Create(&user).Error
}

// DeleteUser deletes the user by the user identifier.
func (user *UserModel) DeleteUser(id uint64) error {
	user.BaseModel.Id = id
	return DB.Self.Delete(&user).Error
}

// Update updates an user account information.
func (user *UserModel) Update() error {
	return DB.Self.Save(user).Error
}

// GetUserByName gets an user by the username.
func (user *UserModel) GetUserByName(username string) (*UserModel, error) {
	d := DB.Self.Where("username = ?", username).First(&user)
	fmt.Println("GetUser-d", d)
	return user, d.Error
}

// GetUserById gets an user by the user id.
func (user *UserModel) GetUserById(id uint64) (*UserModel, error) {
	d := DB.Self.Where("id = ?", id).First(&user)
	fmt.Println("GetUser-d", d)
	return user, d.Error
}

// ListUser List all users
func (user *UserModel) ListUser(offset, limit int) ([]*UserModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	users := make([]*UserModel, 0)
	var count uint64

	if err := DB.Self.Model(&user).Count(&count).Error; err != nil {
		return users, count, err
	}

	if err := DB.Self.Where("").Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}

	return users, count, nil
}

// Compare with the plain text password. Returns true if it's the same as the encrypted one (in the `User` struct).
func (user *UserModel) Compare(pwd string) (err error) {
	err = auth.Compare(user.Password, pwd)
	return
}

// Encrypt the user password.
func (user *UserModel) Encrypt() (err error) {
	user.Password, err = auth.Encrypt(user.Password)
	return
}
