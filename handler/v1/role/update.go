package role

import (
	"strconv"

	. "ant-forum/handler/v1"
	"ant-forum/model"
	"ant-forum/pkg/errno"
	"ant-forum/util"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// @Summary 更新角色数据
// @Description 更新角色数据
// @Tags role
// @Accept  json
// @Produce  json
// @Param id path uint64 true ""
// @Param role body model.RoleModel true ""
// @Success 200 {object} v1.Response "{"code":0,"message":"OK","data":null}"
// @Router /v1/role/{id} [put]
func Update(c *gin.Context) {
	log.Info("Role Update function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	id, _ := strconv.Atoi(c.Param("id"))
	var u model.RoleModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	u.Id = uint64(id)
	if err := u.Update(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
