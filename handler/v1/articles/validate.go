package articles

import (
	"ant-forum/model"
	"fmt"
)

// 检验文章字段是否正确
func ValidateCreateArticle(userId, categoryId, tagId uint64) (valid bool) {
	user, err := model.GetUserById(uint64(userId))

	fmt.Println("1",err)
	if err != nil {
		return false
	}
	category, err := model.GetCategoryById(uint64(categoryId))
	fmt.Println("2",err)
	if err != nil {
		return false
	}
	tag, err := model.GetTagById(uint64(tagId))
	fmt.Println("3",err)
	if err != nil {
		return false
	}

	fmt.Println(user)
	fmt.Println(category)
	fmt.Println(tag)

	if user == nil || category == nil || tag == nil {
		return false
	}


	return true

}
