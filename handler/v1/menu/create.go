package menu

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

// @Summary 创建菜单
// @Description 创建菜单
// @Tags menu
// @Accept  json
// @Produce  json
// @Param tags body menu.CreateRequest true "创建新标签"
// @Success 200 {object} menu.CreateResponse "{"code":0,"message":"OK","data":{"tag_name":"前端"}}"
// @Router /v1/menu [post]
func Create(c *gin.Context) {
	log.Info("Menu Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	u := model.MenuModel{
		Name:   r.Name,
		Method: r.Method,
		Path:   r.Path,
	}

	// Insert the user to the database.
	if err := u.Create(); err != nil {
		fmt.Println(err)
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		Name:   r.Name,
		Method: r.Method,
		Path:   r.Path,
	}

	// Show the user information.
	SendResponse(c, nil, rsp)
}
