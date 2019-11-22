package service

import (
	"ant-forum/model"
)

func ListMenu(offset, limit int) ([]*model.MenuInfo, uint64, error) {
	var infos []*model.MenuInfo
	var m *model.MenuModel
	list, count, err := m.ListMenu(offset, limit)

	if err != nil {
		return nil, count, err
	}
	for _, l := range list {
		infos = append(infos, &model.MenuInfo{
			Id:     l.Id,
			Name:   l.Name,
			Path:   l.Path,
			Method: l.Method,
		})
	}
	return infos, count, nil
}
