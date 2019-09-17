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

func (s *SqlUserStore) SaveGoogleAccount(account *model.GoogleAccount) *model.AppError {
	account.PreSave()

	if _, err := s.Insert("google_account", []model.DbModel{account}); err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_user_store.save_google_account", err.Error(), http.StatusInternalServerError,
			map[string]string{"user_id": account.UserId},
		)
	}
	return nil
}

func (s *SqlUserStore) SaveFacebookAccount(account *model.FacebookAccount) *model.AppError {
	account.PreSave()

	if _, err := s.Insert("facebook_account", []model.DbModel{account}); err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_user_store.save_facebook_account", err.Error(), http.StatusInternalServerError,
			map[string]string{"user_id": account.UserId},
		)
	}
	return nil
}

func (s *SqlUserStore) SaveTwitterAccount(account *model.TwitterAccount) *model.AppError {
	account.PreSave()

	if _, err := s.Insert("twitter_account", []model.DbModel{account}); err != nil {
		return model.NewAppError(
			"store.sqlstore.sql_user_store.save_twitter_account", err.Error(), http.StatusInternalServerError,
			map[string]string{"user_id": account.UserId},
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

func (s *SqlUserStore) GetGoogleAccount(id string) (*model.GoogleAccount, *model.AppError) {
	account := &model.GoogleAccount{}
	sql := "SELECT " + Cols(account) + " FROM google_account WHERE user_id = ?"

	if err := s.GetMaster().QueryRow(sql, id).Scan(account.FieldAddrs()...); err != nil {
		return nil, model.NewAppError(
			"store.sqlstore.sql_user_store.get_google_account", err.Error(), http.StatusInternalServerError,
			map[string]string{"id": id},
		)
	}
	return account, nil
}

func (s *SqlUserStore) GetFacebookAccount(id string) (*model.FacebookAccount, *model.AppError) {
	account := &model.FacebookAccount{}
	sql := "SELECT " + Cols(account) + " FROM facebook_account WHERE user_id = ?"

	if err := s.GetMaster().QueryRow(sql, id).Scan(account.FieldAddrs()...); err != nil {
		return nil, model.NewAppError(
			"store.sqlstore.sql_user_store.get_facebook_account", err.Error(), http.StatusInternalServerError,
			map[string]string{"id": id},
		)
	}
	return account, nil
}

func (s *SqlUserStore) GetTwitterAccount(id string) (*model.FacebookAccount, *model.AppError) {
	account := &model.FacebookAccount{}
	sql := "SELECT " + Cols(account) + " FROM facebook_account WHERE user_id = ?"

	if err := s.GetMaster().QueryRow(sql, id).Scan(account.FieldAddrs()...); err != nil {
		return nil, model.NewAppError(
			"store.sqlstore.sql_user_store.get_facebook_account", err.Error(), http.StatusInternalServerError,
			map[string]string{"id": id},
		)
	}
	return account, nil
}
