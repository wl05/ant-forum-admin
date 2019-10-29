package articles

import (
	. "ant-forum/handler/v1"
	"ant-forum/model"
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

	infos, count, err := service.ListArticles(r.Offset, r.Limit)

	var articleList []*GetArticleInfo
	for _, article := range infos {
		user, uErr := model.GetUserById(uint64(article.UserId))
		category, cErr := model.GetCategoryById(uint64(article.CategoryId))
		tag, tErr := model.GetTagById(uint64(article.TagId))
		if uErr != nil || cErr != nil || tErr != nil {
			SendResponse(c, errno.ErrArticleNotFound, nil)
			return
		}
		articleInfo := &GetArticleInfo{
			Id:           article.Id,
			Title:        article.Title,
			Content:      article.Content,
			CategoryId:   article.CategoryId,
			CategoryName: category.CategoryName,
			TagId:        article.TagId,
			TagName:      tag.TagName,
			UserId:       article.UserId,
			UserName:     user.Username,
			Avatar:       user.Avatar,
			CreatedAt:    article.CreatedAt,
			UpdatedAt:    article.UpdatedAt,
		}

		articleList = append(articleList, articleInfo)

	}

	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		List:       articleList,
	})
}
