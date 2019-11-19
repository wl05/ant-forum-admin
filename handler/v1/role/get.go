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

// @Summary 获取单个角色
// @Description 获取单个角色
// @Tags role
// @Accept  json
// @Produce  json
// @Param id path string true "Id"
// @Success 200 {object} model.RoleInfo "{"code":0,"message":"OK","data":{"id":0,"category_name":"前端"}}"
// @Router /v1/role/{id} [get]
func GetRole(c *gin.Context) {
	log.Info("Role GetRole function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	id, _ := strconv.Atoi(c.Param("id"))
	m := &model.RoleModel{}
	l, err := m.GetRoleById(uint64(id))
	if err != nil {
		SendResponse(c, errno.ErrCategoryNotFound, nil)
		return
	}
	SendResponse(c, nil, &model.RoleInfo{Id: l.Id, Name: l.Name})
}
