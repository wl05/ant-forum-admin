package articles

type CreateRequest struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	CategoryId uint64 `json:"category_id"`
	TagId      uint64 `json:"tag_id"`
	UserId     uint64 `json:"user_id"`
}
