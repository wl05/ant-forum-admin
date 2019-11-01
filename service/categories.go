package service

import (
	"ant-forum/model"
)

func ListCategories(offset, limit int) ([]*model.CategoryInfo, uint64, error) {
	categories, count, err := model.ListCategories(offset, limit)
	if err != nil {
		return nil, count, err
	}
	var infos []*model.CategoryInfo
	for _, category := range categories {
		categoryInfo := &model.CategoryInfo{
			Id:           category.Id,
			CategoryName: category.CategoryName,
		}
		infos = append(infos, categoryInfo)
	}

	return infos, count, nil
}
