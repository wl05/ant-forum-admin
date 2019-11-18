package tags

import "ant-forum/model"

type CreateRequest struct {
	TagName string `json:"tag_name"`
}

type CreateResponse struct {
	TagName string `json:"tag_name"`
}

type ListRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type ListResponse struct {
	TotalCount uint64           `json:"totalCount"`
	List       []*model.TagInfo `json:"list"`
}
