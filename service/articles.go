package service

import (
	"ant-forum/model"
)

func ListArticles(offset, limit int) ([]*model.ArticleInfo, uint64, error) {
	articles, count, err := model.ListArticles(offset, limit)
	if err != nil {
		return nil, count, err
	}

	var articleList []*model.ArticleInfo
	for _, article := range articles {
		user, uErr := model.GetUserById(uint64(article.UserId))
		category, cErr := model.GetCategoryById(uint64(article.CategoryId))
		tag, tErr := model.GetTagById(uint64(article.TagId))
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