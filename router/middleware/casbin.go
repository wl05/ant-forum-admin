package middleware

import (
	"ant-forum/handler/v1"
	"ant-forum/pkg/errno"
	"fmt"
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
	"strings"
)

func CasbinMiddleware(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取请求的URI
		obj := c.Request.URL.Path
		//获取请求方法
		act := strings.ToLower(c.Request.Method)
		//获取用户的角色
		sub := "admin"
		//判断策略中是否存在
		if res, _ := e.Enforce(sub, obj, act); res {
			c.Next()
		} else {
			v1.SendResponse(c, errno.ErrNotPermission, nil)
			c.Abort()
		}
	}
}
