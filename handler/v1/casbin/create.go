package casbin

import (
	. "ant-forum/handler/v1"
	"ant-forum/pkg/errno"
	"ant-forum/util"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// @Summary 创建策略
// @Description 创建策略
// @Tags casbin
// @Accept  json
// @Produce  json
// @Param casbin body casbin.CreateRequest true "创建新标签"
// @Success 200 {object} casbin.CreateResponse "{"code":0,"message":"OK","data":{"category_name":"前端"}}"
// @Router /v1/casbin [post]
func Create(c *gin.Context) {
	log.Info("Casbin Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	err := util.AddCasbin(r.RoleName,r.Source,r.Method)

	if err != nil {
		SendResponse(c, errno.ErrCreateCasbin, nil)
		return
	}

	rsp := CreateResponse{
		RoleName: r.RoleName,
		Source:   r.Source,
		Method:   r.Method,
	}
	SendResponse(c, nil, rsp)
	return
}
