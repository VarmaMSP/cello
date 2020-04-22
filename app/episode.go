package app

import "github.com/varmamsp/cello/model"

func (a *App) GetEpisode(episodeId int64) (*model.Episode, *model.AppError) {
	return a.Store.Episode().Get(episodeId)
}

func (a *App) GetEpisodesByIds(episodeIds []int64) ([]*model.Episode, *model.AppError) {
	if len(episodeIds) == 0 {
		return []*model.Episode{}, nil
	}
	return a.Store.Episode().GetByIds(episodeIds)
}

func (a *App) GetEpisodesByPodcast(podcastId int64, order string, offset, limit int) ([]*model.Episode, *model.AppError) {
	return a.Store.Episode().GetByPodcastPaginated(podcastId, order, offset, limit)
}

func (a *App) GetEpisodesByPodcasts(podcastIds []int64, offset, limit int) ([]*model.Episode, *model.AppError) {
	if len(podcastIds) == 0 {
		return []*model.Episode{}, nil
	}
	return a.Store.Episode().GetByPodcastIdsPaginated(podcastIds, offset, limit)
}
