package sqlstore

import (
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store"
)

type SqlPodcastStore struct {
	SqlStore
}

func (pss *SqlPodcastStore) Save(podcast *model.Podcast) store.StoreChannel {
	return store.Do(func(r *store.StoreResult) {
		pss.Insert([]DbModel{podcast}, "podcast")
	})
}

func NewSqlPodcastStore(store SqlStore) store.PodcastStore {
	return &SqlPodcastStore{store}
}
