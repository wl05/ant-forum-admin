package role

import "ant-forum/model"

type CreateRequest struct {
	Name string `form:"name" json:"name" xml:"name" binding:"required"`
}

type CreateResponse struct {
	Name string `json:"name"`
}

type ListRequest struct {
	Offset int `form:"offset" json:"offset" xml:"offset"`
	Limit  int `form:"limit" json:"limit" xml:"limit"`
}

type ListResponse struct {
	TotalCount uint64            `json:"totalCount"`
	List       []*model.RoleInfo `json:"list"`
}
