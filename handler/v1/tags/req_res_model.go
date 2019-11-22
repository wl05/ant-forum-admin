package tags

import "ant-forum/model"

type CreateRequest struct {
	TagName string `form:"tag_name" json:"tag_name" xml:"tag_name" binding:"required"`

}

type CreateResponse struct {
	TagName string `json:"tag_name"`
}

type ListRequest struct {
	Offset int `form:"offset" json:"offset" xml:"offset"`
	Limit  int `form:"limit" json:"limit" xml:"limit"`
}

type ListResponse struct {
	TotalCount uint64           `json:"totalCount"`
	List       []*model.TagInfo `json:"list"`
}
