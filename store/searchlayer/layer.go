package searchlayer

import (
	"github.com/varmamsp/cello/service/searchengine"
	"github.com/varmamsp/cello/store"
)

type searchStore struct {
	store.Store

	podcast store.PodcastStore
	episode store.EpisodeStore
}

func NewSearchLayer(baseStore store.Store, broker searchengine.Broker) store.Store {
	return &searchStore{
		Store:   baseStore,
		podcast: &searchPodcastStore{PodcastStore: baseStore.Podcast(), se: broker},
		episode: &searchEpisodeStore{EpisodeStore: baseStore.Episode(), se: broker},
	}
}

func (s *searchStore) Podcast() store.PodcastStore {
	return s.podcast
}

func (s *searchStore) Episode() store.EpisodeStore {
	return s.episode
}
