package service

import (
	"ant-forum/model"
)

func ListRole(offset, limit int) ([]*model.RoleInfo, uint64, error) {
	var infos []*model.RoleInfo
	list, count, err := model.ListRole(offset, limit)

	if err != nil {
		return nil, count, err
	}
	for _, l := range list {
		infos = append(infos, &model.RoleInfo{
			Id:   l.Id,
			Name: l.Name,
		})
	}
	return infos, count, nil
}
