package menu

import (
	. "ant-forum/handler/v1"
	"ant-forum/pkg/errno"
	"ant-forum/service"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

// @Summary 获取菜单列表
// @Description 获取菜单列表
// @Tags menu
// @Accept  json
// @Produce  json
// @Param   offset      query    int     false     "Offset"
// @Param   limit      query    int     false      "Limit"
// @Success 200 {object} menu.ListResponse "{"code":0,"message":"OK","data":{"totalCount":1,"list":[{"id":0,"tag_name":"前端"}]}}"
// @Router /v1/menu [get]
func List(c *gin.Context) {
	log.Info("Menu list function called.")
	var r ListRequest
	if err := c.ShouldBindQuery(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	if r.Limit == 0 {
		r.Limit = 10
	}
	infos, count, err := service.ListMenu(r.Offset, r.Limit)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		List:       infos,
	})
}
