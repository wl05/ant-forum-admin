package menu

import (
	"fmt"
	"strconv"

	. "ant-forum/handler/v1"
	"ant-forum/model"
	"ant-forum/pkg/errno"
	"ant-forum/util"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// @Summary 更改菜单
// @Description 更改菜单
// @Tags menu
// @Accept  json
// @Produce  json
// @Param id path uint64 true "菜单数据的数据库id"
// @Param menu body model.MenuModel true "The user info"
// @Success 200 {object} v1.Response "{"code":0,"message":"OK","data":null}"
// @Router /v1/menu/{id} [put]
func Update(c *gin.Context) {
	log.Info("Update function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	id, _ := strconv.Atoi(c.Param("id"))
	var req model.MenuModel
	if err := c.ShouldBindJSON(&req); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	req.Id = uint64(id)
	if err := req.Update(); err != nil {
		fmt.Println(err)
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
