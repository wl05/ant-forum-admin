package service

import (
	"ant-forum/model"
)

func ListUser(offset, limit int) ([]*model.UserInfo, uint64, error) {
	var infos []*model.UserInfo
	var u model.UserModel
	users, count, err := u.ListUser(offset, limit)
	if err != nil {
		return nil, count, err
	}

	for _, user := range users {
		infos = append(infos, &model.UserInfo{
			Id:       user.Id,
			Username: user.Username,
			Avatar:   user.Avatar,
		})
	}

	return infos, count, nil
}
