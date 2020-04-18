package searchlayer

import (
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/elasticsearch_"
	"github.com/varmamsp/cello/store_"
)

type searchStore struct {
	store_.Store

	podcast store_.PodcastStore
	episode store_.EpisodeStore
}

func NewSearchLayer(baseStore store_.Store, broker elasticsearch_.Broker) (store_.Store, *model.AppError) {
	s := &searchStore{Store: baseStore}

	if episodeStore, err := newSearchEpisodeStore(baseStore.Episode(), broker); err != nil {
		return nil, err
	} else {
		s.episode = episodeStore
	}

	if podcastStore, err := newSearchPodcastStore(baseStore.Podcast(), broker); err != nil {
		return nil, err
	} else {
		s.podcast = podcastStore
	}

	return s, nil
}

func (s *searchStore) Podcast() store_.PodcastStore {
	return s.podcast
}

func (s *searchStore) Episode() store_.EpisodeStore {
	return s.episode
}
