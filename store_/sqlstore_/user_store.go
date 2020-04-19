package sqlstore_

import (
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/sqldb"
)

type sqlUserStore struct {
	sqldb.Broker
}

func (s *sqlUserStore) Save(user *model.User) *model.AppError {
	return nil
}

func (s *sqlUserStore) SaveSocialAccount(accountType string, account model.DbModel) *model.AppError {
	return nil
}

func (s *sqlUserStore) Get(userId int64) (*model.User, *model.AppError) {
	return nil, nil
}

func (s *sqlUserStore) GetSocialAccount(accountType, id string) (model.DbModel, *model.AppError) {
	return nil, nil
}
