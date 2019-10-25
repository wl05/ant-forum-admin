package tags

type CreateRequest struct {
	TagName string `json:"tag_name"`
}

type CreateResponse struct {
	TagName string `json:"tag_name"`
}
