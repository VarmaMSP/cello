package sqlstore

import (
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store"
)

type SqlEpisodeStore struct {
	SqlStore
}

func (ess SqlEpisodeStore) SaveAll(episodes []*model.Episode) store.StoreChannel {
	return store.Do(func(result *store.StoreResult) {
		models := make([]DbModel, len(episodes))
		for i := range models {
			models[i] = episodes[i]
		}
		ess.Insert(models, "episode")
	})
}
