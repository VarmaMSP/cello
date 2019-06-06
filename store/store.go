package store

import "github.com/varmamsp/cello/model"

type StoreResult struct {
	Data interface{}
	Err  error
}

type StoreChannel chan StoreResult

func Do(f func(result *StoreResult)) StoreChannel {
	storeChannel := make(StoreChannel, 1)
	go func() {
		storeResult := StoreResult{}
		f(&storeResult)
		storeChannel <- storeResult
		close(storeChannel)
	}()
	return storeChannel
}

type Store interface {
	Podcast() PodcastStore
	Episode() EpisodeStore
}

type PodcastStore interface {
	Save(podcast *model.Podcast) StoreChannel
}

type EpisodeStore interface {
	SaveAll(episodes []*model.Episode) StoreChannel
}
