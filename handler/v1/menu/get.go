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
// @Param id path string true "Id"
// @Success 200 {object} model.CategoryInfo "{"code":0,"message":"OK","data":{"id":0,"category_name":"前端"}}"
// @Router /v1/menu/{id} [get]
func GetMenu(c *gin.Context) {
	log.Info("Category GetCategoryById function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	categoryId, _ := strconv.Atoi(c.Param("id"))
	// Get the category by the `id` from the database.
	category, err := model.GetCategoryById(uint64(categoryId))
	if err != nil {
		SendResponse(c, errno.ErrCategoryNotFound, nil)
		return
	}
	SendResponse(c, nil, &model.CategoryInfo{Id: category.Id, CategoryName: category.CategoryName})
}
