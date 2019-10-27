package tags

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

// @Summary 创建标签
// @Description 创建标签
// @Tags tags
// @Accept  json
// @Produce  json
// @Param tags body tags.CreateRequest true "创建新标签"
// @Success 200 {object} tags.CreateResponse "{"code":0,"message":"OK","data":{"tag_name":"前端"}}"
// @Router /v1/tags [post]
func Create(c *gin.Context) {
	log.Info("Tag Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.TagModel{
		TagName: r.TagName,
	}

	// Validate the data.
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// Insert the user to the database.
	if err := u.Create(); err != nil {
		fmt.Println(err)
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		TagName: r.TagName,
	}

	// Show the user information.
	SendResponse(c, nil, rsp)
}
