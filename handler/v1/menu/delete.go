package menu

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

// @Summary 根据菜单id删除菜单
// @Description 根据菜单id删除菜单
// @Tags menu
// @Accept  json
// @Produce  json
// @Param id path uint64 true "菜单数据的数据库id"
// @Success 200 {object} v1.Response "{"code":0,"message":"OK","data":null}"
// @Router /v1/menu/{id} [delete]
func Delete(c *gin.Context) {
	log.Info("Menu delete function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	id, _ := strconv.Atoi(c.Param("id"))
	m := &model.MenuModel{}
	if err := m.DeleteMenu(uint64(id)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	SendResponse(c, nil, nil)
}
