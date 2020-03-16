package articles

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

// @Summary 根据标签id删除文章
// @Description 根据标签id删除文章
// @Tags articles
// @Accept  json
// @Produce  json
// @Param id path uint64 true "文章的数据库id"
// @Success 200 {object} v1.Response "{"code":0,"message":"OK","data":null}"
// @Router /v1/articles/{id} [delete]
func Delete(c *gin.Context) {
	log.Info("Article Delete function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	articleId, _ := strconv.Atoi(c.Param("id"))
	if err := (&model.ArticleModel{}).DeleteArticle(uint64(articleId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	SendResponse(c, nil, nil)
}
