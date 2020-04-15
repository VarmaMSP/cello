package sqlstore_

import (
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store_"
)

type sqlPodcastStore struct {
	sqlStore
}

func newSqlPodcastStore(store sqlStore) store_.PodcastStore {
	return &sqlPodcastStore{store}
}

func (s *sqlPodcastStore) Save(podcast *model.Podcast) *model.AppError {
	panic("not implemented") // TODO: Implement
}

func (s *sqlPodcastStore) Get(podcastId int64) (*model.Podcast, *model.AppError) {
	panic("not implemented") // TODO: Implement
}

func (s *sqlPodcastStore) GetAllPaginated(lastId int64, limit int) ([]*model.Podcast, *model.AppError) {
	panic("not implemented") // TODO: Implement
}

func (s *sqlPodcastStore) GetByIds(podcastIds []int64) ([]*model.Podcast, *model.AppError) {
	panic("not implemented") // TODO: Implement
}

func (s *sqlPodcastStore) GetSubscriptions(userId int64) ([]*model.Podcast, *model.AppError) {
	panic("not implemented") // TODO: Implement
}

func (s *sqlPodcastStore) Update(old *model.Podcast, new *model.Podcast) *model.AppError {
	panic("not implemented") // TODO: Implement
}
