package role

import "ant-forum/model"

type CreateRequest struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	Method string `json:"method"`
}

type CreateResponse struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	Method string `json:"method"`
}

type ListRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type ListResponse struct {
	TotalCount uint64            `json:"totalCount"`
	List       []*model.RoleInfo `json:"list"`
}
