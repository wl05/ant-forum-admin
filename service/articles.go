package service

import (
	"ant-forum/model"
	"sync"
)

func ListArticles(offset, limit int) ([]*model.ArticleInfo, uint64, error) {
	infos := make([]*model.ArticleInfo, 0)
	articles, count, err := model.ListArticles(offset, limit)
	if err != nil {
		return nil, count, err
	}

	ids := []uint64{}
	for _, tag := range articles {
		ids = append(ids, tag.Id)
	}

	wg := sync.WaitGroup{}
	articleList := model.ArticlesList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*model.ArticleInfo, len(articles)),
	}

	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	// Improve query efficiency in parallel
	for _, u := range articles {
		wg.Add(1)
		go func(u *model.ArticleModel) {
			defer wg.Done()

			if err != nil {
				errChan <- err
				return
			}

			articleList.Lock.Lock()
			defer articleList.Lock.Unlock()
			articleList.IdMap[u.Id] = &model.ArticleInfo{
				Id:         u.Id,
				Title:      u.Title,
				Content:    u.Content,
				CategoryId: u.CategoryId,
				TagId:      u.TagId,
				UserId:     u.UserId,
				CreatedAt:  u.CreatedAt,
				UpdatedAt:  u.UpdatedAt,
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
		infos = append(infos, articleList.IdMap[id])
	}

	return infos, count, nil
}
