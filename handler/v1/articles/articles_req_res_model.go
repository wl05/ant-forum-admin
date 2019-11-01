package articles

import "ant-forum/model"

type CreateRequest struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	CategoryId uint64 `json:"category_id"`
	TagId      uint64 `json:"tag_id"`
	UserId     uint64 `json:"user_id"`
}
//
//type GetArticleInfo struct {
//	Id           uint64    `json:"id"`
//	Title        string    `json:"title"`
//	Content      string    `json:"content"`
//	CategoryId   uint64    `json:"category_id"`
//	CategoryName string    `json:"category_name"`
//	TagId        uint64    `json:"tag_id"`
//	TagName      string    `json:"tag_name"`
//	UserId       uint64    `json:"user_id"`
//	UserName     string    `json:"username"`
//	Avatar       string    `json:"avatar"`
//	CreatedAt    time.Time `json:"created_at"`
//	UpdatedAt    time.Time `json:"update_at"`
//}

type ListRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type ListResponse struct {
	TotalCount uint64            `json:"totalCount"`
	List       []*model.ArticleInfo `json:"list"`
}
