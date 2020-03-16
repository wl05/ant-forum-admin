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

// @Summary 获取单个菜单
// @Description 获取单个菜单
// @Tags menu
// @Accept  json
// @Produce  json
// @Param id path uint64 true "菜单数据的数据库id"
// @Success 200 {object} model.MenuInfo "{"code":0,"message":"OK","data":{"id":0,"category_name":"前端"}}"
// @Router /v1/menu/{id} [get]
func GetMenu(c *gin.Context) {
	log.Info("Menu GetMenu function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	id, _ := strconv.Atoi(c.Param("id"))
	l, err := (model.MenuModel{}).GetMenuById(uint64(id))
	if err != nil {
		SendResponse(c, errno.ErrMenuGet, nil)
		return
	}
	SendResponse(c, nil, &model.MenuInfo{Id: l.Id, Name: l.Name, Path: l.Path, Method: l.Method})
}
