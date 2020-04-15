package sqlstore_

import (
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store_"
)

type sqlEpisodeStore struct {
	sqlStore
}

func newSqlEpisodeStore(store sqlStore) store_.EpisodeStore {
	return &sqlEpisodeStore{store}
}

func (s *sqlEpisodeStore) Save(episode *model.Episode) *model.AppError {
	panic("not implemented") // TODO: Implement
}

func (s *sqlEpisodeStore) Get(episodeId int64) (*model.Episode, *model.AppError) {
	panic("not implemented") // TODO: Implement
}

func (s *sqlEpisodeStore) GetAllPaginated(lastId int64, limit int) ([]*model.Episode, *model.AppError) {
	panic("not implemented") // TODO: Implement
}

func (s *sqlEpisodeStore) GetByIds(episodeIds []int64) ([]*model.Episode, *model.AppError) {
	panic("not implemented") // TODO: Implement
}

func (s *sqlEpisodeStore) GetByPodcast(podcastId int64) ([]*model.Episode, *model.AppError) {
	panic("not implemented") // TODO: Implement
}

func (s *sqlEpisodeStore) GetByPodcastPaginated(podcastId int64, order string, offset int, limit int) ([]*model.Episode, *model.AppError) {
	panic("not implemented") // TODO: Implement
}

func (s *sqlEpisodeStore) GetByPodcastIdsPaginated(podcastIds []int64, offset int, limit int) ([]*model.Episode, *model.AppError) {
	panic("not implemented") // TODO: Implement
}

func (s *sqlEpisodeStore) GetByPlaylistPaginated(playlistId int64, offset int, limit int) ([]*model.Episode, *model.AppError) {
	panic("not implemented") // TODO: Implement
}

func (s *sqlEpisodeStore) Block(episodeIds []int64) *model.AppError {
	panic("not implemented") // TODO: Implement
}
