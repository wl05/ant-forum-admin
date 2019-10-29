package model

import (
	"fmt"
	"sync"

	"ant-forum/pkg/auth"
	"ant-forum/pkg/constvar"

	validator "gopkg.in/go-playground/validator.v9"
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

type UserList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*UserInfo
}

func (c *UserModel) TableName() string {
	return "users"
}

// Create creates a new user account.
func (u *UserModel) Create() error {
	return DB.Self.Create(&u).Error
}

// DeleteUser deletes the user by the user identifier.
func DeleteUser(id uint64) error {
	user := UserModel{}
	user.BaseModel.Id = id
	return DB.Self.Delete(&user).Error
}

// Update updates an user account information.
func (u *UserModel) Update() error {
	return DB.Self.Save(u).Error
}

// GetUserByName gets an user by the username.
func GetUserByName(username string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Where("username = ?", username).First(&u)
	fmt.Println("GetUser-d", d)
	return u, d.Error
}

// GetUserById gets an user by the user id.
func GetUserById(id uint64) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Where("id = ?", id).First(&u)
	fmt.Println("GetUser-d", d)
	return u, d.Error
}

// ListUser List all users
func ListUser(offset, limit int) ([]*UserModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	users := make([]*UserModel, 0)
	var count uint64

	if err := DB.Self.Model(&UserModel{}).Count(&count).Error; err != nil {
		return users, count, err
	}

	if err := DB.Self.Where("").Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}

	return users, count, nil
}

// Compare with the plain text password. Returns true if it's the same as the encrypted one (in the `User` struct).
func (u *UserModel) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)
	return
}

// Encrypt the user password.
func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

// Validate the fields.
func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
