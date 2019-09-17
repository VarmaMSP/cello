package sqlstore

import (
	"net/http"

	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store"
)

type SqlUserStore struct {
	SqlStore
}

func NewSqlUserStore(store SqlStore) store.UserStore {
	return &SqlUserStore{store}
}

func (s *SqlUserStore) Save(user *model.User) *model.AppError {
	user.PreSave()

	if _, err := s.Insert("user", []model.DbModel{user}); err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_user_store.save", err.Error(), http.StatusInternalServerError,
			map[string]string{"name": user.Name},
		)
	}
	return nil
}

func (s *SqlUserStore) SaveSocialAccount(accountType string, account model.DbModel) *model.AppError {
	if accountType != "google" && accountType != "facebook" && accountType != "twitter" {
		return nil
	}

	account.PreSave()

	if _, err := s.Insert(accountType+"_account", []model.DbModel{account}); err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_user_store.save_google_account", err.Error(), http.StatusInternalServerError,
			map[string]string{"account_type": accountType},
		)
	}
	return nil
}

func (s *SqlUserStore) Get(userId string) (*model.User, *model.AppError) {
	user := &model.User{}
	sql := "SELECT " + Cols(user) + " FROM user WHERE id = ?"

	if err := s.GetMaster().QueryRow(sql, userId).Scan(user.FieldAddrs()...); err != nil {
		return nil, model.NewAppError(
			"store.sqlstore.sql_user_store.get", err.Error(), http.StatusInternalServerError,
			map[string]string{"user_id": userId},
		)
	}
	return user, nil
}

func (s *SqlUserStore) GetSocialAccount(accountType, id string) (model.DbModel, *model.AppError) {
	var account model.DbModel
	if accountType == "google" {
		account = &model.GoogleAccount{}
	} else if accountType == "facebook" {
		account = &model.FacebookAccount{}
	} else if accountType == "twitter" {
		account = &model.TwitterAccount{}
	} else {
		return nil, nil
	}

	tableName := accountType + "_account"
	sql := "SELECT " + Cols(account) + " FROM " + tableName + " WHERE id = ?"

	if err := s.GetMaster().QueryRow(sql, id).Scan(account.FieldAddrs()...); err != nil {
		return nil, model.NewAppError(
			"store.sqlstore.sql_user_store.get_social_account", err.Error(), http.StatusInternalServerError,
			map[string]string{"account_type": accountType, "id": id},
		)
	}
	return account, nil
}
