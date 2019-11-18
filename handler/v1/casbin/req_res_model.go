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

type ListResponse struct {
	List []*CreateResponse `json:"list"`
}
