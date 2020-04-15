package searchlayer

import (
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/searchengine"
	"github.com/varmamsp/cello/store_"
)

type searchPodcastStore struct {
	store_.PodcastStore
	podcastIndex searchengine.PodcastIndex
}

func (s *searchPodcastStore) Save(podcast *model.Podcast) *model.AppError {
	if err := s.podcastIndex.Save(podcast); err != nil {
		return err
	}

	return s.PodcastStore.Save(podcast)
}
