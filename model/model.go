package model

import (
	"sync"
	"time"
)

type BaseModel struct {
	Id        uint64     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	Created_at time.Time  `gorm:"column:createdAt" json:"-"`
	Updated_at time.Time  `gorm:"column:updatedAt" json:"-"`
	Deleted_at *time.Time `gorm:"column:deletedAt" sql:"index" json:"-"`
}

type UserInfo struct {
	Id        uint64 `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Avatar     string `json:"avatar"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"Updated_at"`
	Deleted_at time.Time `json:"Deleted_at"`
}

type UserList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*UserInfo
}

// Token represents a JSON web token.
type Token struct {
	Token string `json:"token"`
}
