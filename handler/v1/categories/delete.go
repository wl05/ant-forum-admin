package categories

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

// @Summary 根据标签id删除分类
// @Description 根据标签id删除分类
// @Tags categories
// @Accept  json
// @Produce  json
// @Param id path uint64 true "标签数据的数据库id"
// @Success 200 {object} v1.Response "{"code":0,"message":"OK","data":null}"
// @Router /v1/categories/{id} [delete]
func Delete(c *gin.Context) {
	log.Info("Category delete function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	tagId, _ := strconv.Atoi(c.Param("id"))
	if err := (model.CategoriesModel{}).DeleteCategory(uint64(tagId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	SendResponse(c, nil, nil)
}
