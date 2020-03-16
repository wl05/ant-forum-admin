package tags

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
// @Tags tags
// @Accept  json
// @Produce  json
// @Param id path uint64 true "角色数据的数据库id"
// @Param tags body model.TagModel true "role"
// @Success 200 {object} v1.Response "{"code":0,"message":"OK","data":null}"
// @Router /v1/role/{id} [put]
func Update(c *gin.Context) {
	log.Info("Tag Update function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	id, _ := strconv.Atoi(c.Param("id"))
	var tag model.TagModel
	if err := c.ShouldBindJSON(&tag); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	tag.Id = uint64(id)
	if err := tag.Update(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
