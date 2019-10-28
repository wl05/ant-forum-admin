package categories

import "ant-forum/model"

type CreateRequest struct {
	CategoryName string `json:"category_name"`
}

type CreateResponse struct {
	CategoryName string `json:"category_name"`
}

type ListRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type ListResponse struct {
	TotalCount uint64                `json:"totalCount"`
	List       []*model.CategoryInfo `json:"list"`
}
