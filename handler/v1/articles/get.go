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

// @Summary 用文章id获取单篇信息
// @Description 用文章id获取单篇文章信息
// @Tags articles
// @Accept  json
// @Produce  json
// @Param id path string true "Id"
// @Success 200 {object} articles.GetArticleInfo "{"code":0,"message":"OK","data":{}}"
// @Router /v1/articles/{id} [get]
func GetArticleById(c *gin.Context) {
	log.Info("Article GetArticleById function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	articleId, _ := strconv.Atoi(c.Param("id"))
	// Get the article by the `id` from the database.
	article, err := model.GetArticleById(uint64(articleId))
	if err != nil {
		SendResponse(c, errno.ErrArticleNotFound, nil)
		return
	}
	user, uErr := model.GetUserById(uint64(article.UserId))
	category, cErr := model.GetCategoryById(uint64(article.CategoryId))
	tag, tErr := model.GetTagById(uint64(article.TagId))

	if uErr != nil || cErr != nil || tErr != nil {
		SendResponse(c, errno.ErrArticleNotFound, nil)
		return
	}
	articleInfo := &model.ArticleInfo{
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

	SendResponse(c, nil, articleInfo)
}
