package searchlayer

import (
	"github.com/varmamsp/cello/service/searchengine"
	"github.com/varmamsp/cello/store_"
)

type searchStore struct {
	store_.Store

	podcast store_.PodcastStore
	episode store_.EpisodeStore
}

func NewSearchLayer(baseStore store_.Store, broker searchengine.Broker) store_.Store {
	return &searchStore{
		Store:   baseStore,
		podcast: &searchPodcastStore{PodcastStore: baseStore.Podcast(), search: broker},
		episode: &searchEpisodeStore{EpisodeStore: baseStore.Episode(), search: broker},
	}
}

func (s *searchStore) Podcast() store_.PodcastStore {
	return s.podcast
}

func (s *searchStore) Episode() store_.EpisodeStore {
	return s.episode
}
