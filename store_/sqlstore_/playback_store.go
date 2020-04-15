package sqlstore_

import (
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store_"
)

type sqlPlaybackStore struct {
	sqlStore
}

func newSqlPlaybackStore(store sqlStore) store_.PlaybackStore {
	return &sqlPlaybackStore{store}
}

func (s *sqlPlaybackStore) Save(playback *model.Playback) *model.AppError {
	panic("not implemented") // TODO: Implement
}

func (s *sqlPlaybackStore) Upsert(playback *model.Playback) *model.AppError {
	panic("not implemented") // TODO: Implement
}

func (s *sqlPlaybackStore) GetByUserPaginated(userId int64, offset int, limit int) ([]*model.Playback, *model.AppError) {
	panic("not implemented") // TODO: Implement
}

func (s *sqlPlaybackStore) GetByUserByEpisodes(userId int64, episodeIds []int64) ([]*model.Playback, *model.AppError) {
	panic("not implemented") // TODO: Implement
}

func (s *sqlPlaybackStore) Update(playback *model.Playback) *model.AppError {
	panic("not implemented") // TODO: Implement
}
