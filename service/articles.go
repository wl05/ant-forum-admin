package service

import (
	"ant-forum/model"
)

func ListArticles(offset, limit int) ([]*model.ArticleInfo, uint64, error) {
	var article *model.ArticleModel
	articles, count, err := article.ListArticles(offset, limit)
	if err != nil {
		return nil, count, err
	}

	var articleList []*model.ArticleInfo
	for _, article := range articles {
		var u *model.UserModel
		var t *model.TagModel
		var g *model.CategoriesModel
		user, uErr := u.GetUserById(uint64(article.UserId))
		category, cErr := g.GetCategoryById(uint64(article.CategoryId))
		tag, tErr := t.GetTagById(uint64(article.TagId))
		if uErr != nil || cErr!= nil || tErr != nil{
			return nil, count, err
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
		articleList = append(articleList, articleInfo)
	}
	return articleList, count, nil
}