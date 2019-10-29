package articles

import (
	"ant-forum/model"
)

// 检验文章字段是否正确
func ValidateCreateArticle(userId, categoryId, tagId uint64) (valid bool) {
	user, err := model.GetUserById(uint64(userId))
	if err != nil {
		return false
	}
	category, err := model.GetCategoryById(uint64(categoryId))
	if err != nil {
		return false
	}
	tag, err := model.GetTagById(uint64(tagId))
	if err != nil {
		return false
	}

	if user == nil || category == nil || tag == nil {
		return false
	}

	return true

}
