package menu

import "ant-forum/model"

type CreateRequest struct {
	Name   string `form:"name" json:"name" xml:"name"  binding:"required"`
	Path   string `form:"path" json:"path" xml:"path" binding:"required"`
	Method string `form:"method" json:"method" xml:"method" binding:"required"`
}

type CreateResponse struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	Method string `json:"method"`
}

type ListRequest struct {
	Offset int `form:"offset" json:"offset" xml:"offset"`
	Limit  int `form:"limit" json:"limit" xml:"limit"`
}

type ListResponse struct {
	TotalCount uint64            `json:"totalCount"`
	List       []*model.MenuInfo `json:"list"`
}

type Get struct {
	ID string `uri:"id" binding:"required,uuid"`
}
