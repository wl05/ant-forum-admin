package casbin

import (
	. "ant-forum/handler/v1"
	"ant-forum/pkg/errno"
	"ant-forum/util"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

// @Summary 获取casbin列表
// @Description 获取casbin列表
// @Tags casbin
// @Accept  json
// @Produce  json
// @Success 200 {object} casbin.ListResponse "{"code":0,"message":"OK","data":{"totalCount":1,"list":[{"id":0,"category_name":"前端"}]}}"
// @Router /v1/casbin [get]
func List(c *gin.Context) {
	log.Info("casbin list function called.")
	var r ListResponse
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	policy := util.GetPolicy()
	SendResponse(c, nil, policy)
	return
}
