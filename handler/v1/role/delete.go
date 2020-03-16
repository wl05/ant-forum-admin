package role

import (
	. "ant-forum/handler/v1"
	"ant-forum/model"
	"ant-forum/pkg/errno"
	"ant-forum/util"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"strconv"
)

// @Summary 根据角色id删除角色
// @Description 根据角色id删除角色
// @Tags role
// @Accept  json
// @Produce  json
// @Param id path uint64 true "角色数据的数据库id"
// @Success 200 {object} v1.Response "{"code":0,"message":"OK","data":null}"
// @Router /v1/role/{id} [delete]
func Delete(c *gin.Context) {
	log.Info("Role delete function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	id, _ := strconv.Atoi(c.Param("id"))
	if err := (model.RoleModel{}).DeleteRole(uint64(id)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	SendResponse(c, nil, nil)
}
