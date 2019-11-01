package articles

import (
	. "ant-forum/handler/v1"
	"ant-forum/pkg/errno"
	"ant-forum/service"
	"ant-forum/util"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// @Summary 获取文章列表
// @Description 获取文章列表
// @Tags articles
// @Accept  json
// @Produce  json
// @Param   offset      query    int     true     "Offset"
// @Param   limit      query    int     true      "Limit"
// @Success 200 {object} articles.ListResponse "{"code":0,"message":"OK","data":{"totalCount":1,"list":[]"
// @Router /v1/articles [get]
func List(c *gin.Context) {
	log.Info("Articles List function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	articleList, count, err := service.ListArticles(r.Offset, r.Limit)

	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		List:       articleList,
	})
}
