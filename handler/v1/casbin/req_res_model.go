package casbin

type CreateRequest struct {
	RoleName string `json:"role_name"`
	Source   string `json:"source"`
	Method   string `json:"method"`
}

type CreateResponse struct {
	RoleName string `json:"role_name"`
	Source   string `json:"source"`
	Method   string `json:"method"`
}

type ListRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type ListResponse struct {
	TotalCount uint64           `json:"totalCount"`
	List       []*CreateRequest `json:"list"`
}
