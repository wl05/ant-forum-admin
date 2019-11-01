package service

import (
	"ant-forum/model"
)

func ListTags(offset, limit int) ([]*model.TagInfo, uint64, error) {
	var infos []*model.TagInfo
	tags, count, err := model.ListTags(offset, limit)

	if err != nil {
		return nil, count, err
	}
	for _, tag := range tags {
		infos = append(infos, &model.TagInfo{
			Id:      tag.Id,
			TagName: tag.TagName,
		})
	}
	return infos, count, nil
}
