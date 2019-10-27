package service

import (
	"sync"

	"ant-forum/model"
)

func ListTags(offset, limit int) ([]*model.TagInfo, uint64, error) {
	infos := make([]*model.TagInfo, 0)
	tags, count, err := model.ListTags(offset, limit)
	if err != nil {
		return nil, count, err
	}

	ids := []uint64{}
	for _, tag := range tags {
		ids = append(ids, tag.Id)
	}

	wg := sync.WaitGroup{}
	tagList := model.TagsList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*model.TagInfo, len(tags)),
	}

	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	// Improve query efficiency in parallel
	for _, u := range tags {
		wg.Add(1)
		go func(u *model.TagModel) {
			defer wg.Done()

			if err != nil {
				errChan <- err
				return
			}

			tagList.Lock.Lock()
			defer tagList.Lock.Unlock()
			tagList.IdMap[u.Id] = &model.TagInfo{
				Id:      u.Id,
				TagName: u.TagName,
			}
		}(u)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return nil, count, err
	}

	for _, id := range ids {
		infos = append(infos, tagList.IdMap[id])
	}

	return infos, count, nil
}
