package categories

import (
	. "ant-forum/handler/v1"
	"ant-forum/model"
	"ant-forum/pkg/errno"
	"ant-forum/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// @Summary 创建分类
// @Description 创建分类
// @Tags categories
// @Accept  json
// @Produce  json
// @Param tags body categories.CreateRequest true "创建新标签"
// @Success 200 {object} categories.CreateResponse "{"code":0,"message":"OK","data":{"category_name":"前端"}}"
// @Router /v1/categories [post]
func Create(c *gin.Context) {
	log.Info("Category Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	u := model.CategoriesModel{
		CategoryName: r.CategoryName,
	}

	// Insert the user to the database.
	if err := u.Create(); err != nil {
		fmt.Println(err)
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		CategoryName: r.CategoryName,
	}

	// Show the user information.
	SendResponse(c, nil, rsp)
}
