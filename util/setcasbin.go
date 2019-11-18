// Package role provides ...
package util

import (
	"ant-forum/model"
	"fmt"
	"github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
	_ "github.com/go-sql-driver/mysql"
)

type PolicyItem struct {
	Sub string
	Obj string
	Act string
}
// 初始化策略
func InitPolicy(e *casbin.Enforcer) {
	var initPolicy []PolicyItem = []PolicyItem{
		PolicyItem{"admin", "/v1/auth/info", "get"},
		PolicyItem{"admin", "/v1/user", "get"},
		PolicyItem{"admin", "/v1/user/:id", "get"},
		PolicyItem{"admin", "/v1/user", "post"},
		PolicyItem{"admin", "/v1/user/:id", "delete"},
		PolicyItem{"admin", "/v1/casbin", "post"},
		PolicyItem{"admin", "/v1/casbin", "get"},

	}
	for _, p := range initPolicy {
		_, _ = e.AddPolicy(p.Sub, p.Obj, p.Act)
	}

}

// 持久化到数据库
func Casbin() *casbin.Enforcer {
	a, _ := gormadapter.NewAdapterByDB(model.DB.Self)
	e, _ := casbin.NewEnforcer("conf/auth_model.conf", a)

	//從DB加載策略
	err := e.LoadPolicy()
	if err != nil {
		fmt.Println("加载Casbin策略失败")
	} else {
		fmt.Println("加载Casbin策略成功")
	}

	InitPolicy(e)
	return e
}

// 添加权限
func AddCasbin(sub string, obj string, act string) error {
	e := Casbin()
	_, _ = e.AddPolicy(sub, obj, act)
	return e.SavePolicy()
}

// 获取权限列表
func GetPolicy()  [][]string{
	e := Casbin()
	return e.GetPolicy()
}

//GetPolicy


