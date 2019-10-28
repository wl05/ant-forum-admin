package articles

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

// @Summary 创建文章
// @Description 创建文章
// @Tags articles
// @Accept  json
// @Produce  json
// @Param articles body articles.CreateRequest true "创建文章"
// @Success 200 {object} v1.Response "{"code":0,"message":"OK","data":null}"
// @Router /v1/articles [post]
func Create(c *gin.Context) {
	log.Info("Article Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.ArticleModel{
		Title:      r.Title,
		Content:    r.Content,
		CategoryId: r.CategoryId,
		TagId:      r.TagId,
		UserId:     r.UserId,
	}

	// 检验必要字段
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	// 检验字段的合法性
	if valid := ValidateCreateArticle(u.UserId, u.CategoryId, u.TagId); !valid {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// 写入数据库
	if err := u.Create(); err != nil {
		fmt.Println(err)
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
