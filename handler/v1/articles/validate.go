package articles

import (
	"ant-forum/model"
)

// 检验文章字段是否正确
func ValidateCreateArticle(userId, categoryId, tagId uint64) (valid bool) {
	var u *model.UserModel
	user, err := u.GetUserById(uint64(userId))
	if err != nil {
		return false
	}
	var g *model.CategoriesModel
	category, err := g.GetCategoryById(uint64(categoryId))
	if err != nil {
		return false
	}
	var t *model.TagModel
	tag, err := t.GetTagById(uint64(tagId))
	if err != nil {
		return false
	}

	if user == nil || category == nil || tag == nil {
		return false
	}

	return true

}
