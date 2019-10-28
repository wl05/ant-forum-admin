package tags

import (
	. "ant-forum/handler/v1"
	"ant-forum/model"
	"ant-forum/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Summary 根据标签id删除标签
// @Description 根据标签id删除标签
// @Tags tags
// @Accept  json
// @Produce  json
// @Param id path uint64 true "标签数据的数据库id"
// @Success 200 {object} v1.Response "{"code":0,"message":"OK","data":null}"
// @Router /v1/tags/{id} [delete]
func Delete(c *gin.Context) {
	tagId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteTag(uint64(tagId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	SendResponse(c, nil, nil)
}
