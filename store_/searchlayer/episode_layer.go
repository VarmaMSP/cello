package searchlayer

import (
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/searchengine"
	"github.com/varmamsp/cello/store_"
)

type searchEpisodeStore struct {
	store_.EpisodeStore
	episodeIndex searchengine.EpisodeIndex
}

func (s *searchEpisodeStore) Save(episode *model.Episode) *model.AppError {
	if err := s.episodeIndex.Save(episode); err != nil {
		return nil
	}

	return s.EpisodeStore.Save(episode)
}
