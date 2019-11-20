package role

import (
	. "ant-forum/handler/v1"
	"ant-forum/pkg/errno"
	"ant-forum/model"
	"ant-forum/util"

	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// @Summary 创建角色
// @Description 创建角色
// @Tags role
// @Accept  json
// @Produce  json
// @Param role body role.CreateRequest true "创建角色"
// @Success 200 {object} role.CreateResponse "{"code":0,"message":"OK","data":{"tag_name":"前端"}}"
// @Router /v1/role [post]
func Create(c *gin.Context) {
	log.Info("Role Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	u := model.RoleModel{
		Name: r.Name,
	}

	if err := u.Create(); err != nil {
		fmt.Println(err)
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		Name: r.Name,
	}

	// Show the user information.
	SendResponse(c, nil, rsp)
}
