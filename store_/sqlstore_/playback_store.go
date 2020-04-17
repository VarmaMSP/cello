package sqlstore_

import (
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/sqldb"
	"github.com/varmamsp/cello/store_"
)

type sqlPlaybackStore struct {
	sqldb.Broker
}

func newSqlPlaybackStore(broker sqldb.Broker) store_.PlaybackStore {
	return &sqlPlaybackStore{broker}
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
