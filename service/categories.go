package service

import (
	"sync"

	"ant-forum/model"
)

func ListCategories(offset, limit int) ([]*model.CategoryInfo, uint64, error) {
	infos := make([]*model.CategoryInfo, 0)
	categories, count, err := model.ListCategories(offset, limit)
	if err != nil {
		return nil, count, err
	}

	ids := []uint64{}
	for _, tag := range categories {
		ids = append(ids, tag.Id)
	}

	wg := sync.WaitGroup{}
	categoryList := model.CategoriesList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*model.CategoryInfo, len(categories)),
	}

	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	// Improve query efficiency in parallel
	for _, u := range categories {
		wg.Add(1)
		go func(u *model.CategoriesModel) {
			defer wg.Done()

			if err != nil {
				errChan <- err
				return
			}

			categoryList.Lock.Lock()
			defer categoryList.Lock.Unlock()
			categoryList.IdMap[u.Id] = &model.CategoryInfo{
				Id:           u.Id,
				CategoryName: u.CategoryName,
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
		infos = append(infos, categoryList.IdMap[id])
	}

	return infos, count, nil
}
