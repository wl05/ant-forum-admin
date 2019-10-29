package tags

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

// @Summary 用标签id获取单个标签信息
// @Description 用标签id获取单个标签信息
// @Tags tags
// @Accept  json
// @Produce  json
// @Param id path string true "Id"
// @Success 200 {object} model.TagInfo "{"code":0,"message":"OK","data":{"id":0,"tag_name":"前端"}}"
// @Router /v1/tags/{id} [get]
func GetTagById(c *gin.Context) {
	log.Info("Tag delete function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	tagId, _ := strconv.Atoi(c.Param("id"))
	// Get the tag by the `id` from the database.
	tag, err := model.GetTagById(uint64(tagId))
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	SendResponse(c, nil, &model.TagInfo{Id: tag.Id, TagName: tag.TagName})
}
