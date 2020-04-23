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

func (a *App) GetEpisodesFromPodcast(podcastId int64, order string, offset, limit int) ([]*model.Episode, *model.AppError) {
	return a.Store.Episode().GetByPodcastPaginated(podcastId, order, offset, limit)
}

func (a *App) GetEpisodesFromPodcasts(podcasts []*model.Podcast, offset, limit int) ([]*model.Episode, *model.AppError) {
	if len(podcasts) == 0 {
		return []*model.Episode{}, nil
	}

	podcastIds := make([]int64, len(podcasts))
	for i, podcast := range podcasts {
		podcastIds[i] = podcast.Id
	}
	return a.Store.Episode().GetByPodcastIdsPaginated(podcastIds, offset, limit)
}

func (a *App) GetRecentlyPlayedEpisodes(userId int64, offset, limit int) ([]*model.Episode, *model.AppError) {
	playbacks, err := a.GetPlaybacks(userId, offset, limit)
	if err != nil {
		return nil, err
	}

	episodeIds := make([]int64, len(playbacks))
	for i, playback := range playbacks {
		episodeIds[i] = playback.EpisodeId
	}
	episodes, err := a.GetEpisodesByIds(episodeIds)
	if err != nil {
		return nil, err
	}

	playbackByEpisodeId := map[int64]*model.Playback{}
	for _, playback := range playbacks {
		playbackByEpisodeId[playback.EpisodeId] = playback
	}

	for _, episode := range episodes {
		if playback, ok := playbackByEpisodeId[episode.Id]; ok {
			episode.Progress = playback.CurrentProgress
			episode.LastPlayedAt = playback.LastPlayedAt
		}
	}

	return episodes, nil
}

func (a *App) GetPlaylistEpisodes(playlist *model.Playlist) ([]*model.Episode, *model.AppError) {
	episodeIds := make([]int64, len(playlist.Members))
	for i, member := range playlist.Members {
		episodeIds[i] = member.EpisodeId
	}

	return a.GetEpisodesByIds(episodeIds)
}
