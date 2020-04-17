package sqlstore_

import (
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/sqldb"
	"github.com/varmamsp/cello/store_"
)

type sqlPodcastStore struct {
	sqldb.Broker
}

func newSqlPodcastStore(broker sqldb.Broker) store_.PodcastStore {
	return &sqlPodcastStore{broker}
}

func (s *sqlPodcastStore) Save(podcast *model.Podcast) *model.AppError {
	panic("")
}

func (s *sqlPodcastStore) Get(podcastId int64) (*model.Podcast, *model.AppError) {
	panic("")
}

func (s *sqlPodcastStore) GetAllPaginated(lastId int64, limit int) (res []*model.Podcast, appE *model.AppError) {
	panic("")
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

func (s *sqlPodcastStore) Search(query string) ([]*model.Podcast, *model.AppError) {
	panic("not implemented") // TODO: Implement
}
