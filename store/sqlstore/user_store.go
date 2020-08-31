package sqlstore

import (
	"github.com/leporo/sqlf"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/sqldb"
)

type sqlUserStore struct {
	sqldb.Broker
}

func (s *sqlUserStore) Save(user *model.User) *model.AppError {
	user.PreSave()

	res, err := s.Insert_("user", user)
	if err != nil {
		return model.New500Error("sql_store.sql_user_store.save", err.Error(), nil)
	}
	user.Id, _ = res.LastInsertId()
	return nil
}

func (s *sqlUserStore) SaveGuestAccount(account *model.GuestAccount) *model.AppError {
	account.PreSave()

	if _, err := s.Insert("guest_account", account); err != nil {
		return model.New500Error("sql_store.sql_user_store.save_social_account", err.Error(), nil)
	}
	return nil
}

func (s *sqlUserStore) SaveSocialAccount(accountType string, account model.DbModel) *model.AppError {
	if accountType != "google" && accountType != "facebook" && accountType != "twitter" {
		return model.New400Error("sql_store.sql_user_store.save_social_account", "Wrong account type", nil)
	}
	account.PreSave()

	table := accountType + "_account"
	if _, err := s.Insert(table, account); err != nil {
		return model.New500Error("sql_store.sql_user_store.save_social_account", err.Error(), nil)
	}
	return nil
}

func (s *sqlUserStore) Get(userId int64) (*model.User, *model.AppError) {
	query := sqlf.
		Select("*").
		From("podcast").
		Where("id = ?", userId)

	var user model.User
	if err := s.QueryRow(&user, query); err != nil {
		return nil, model.New500Error("sql_store.sql_user_store.get", err.Error(), nil)
	}
	return &user, nil
}

func (s *sqlUserStore) GetSocialAccount(accountType, id string) (model.DbModel, *model.AppError) {
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

	query := sqlf.
		Select("*").
		From(accountType+"_account").
		Where("id = ?", id)

	if err := s.QueryRow_(account.FieldAddrs(), query.String(), query.Args()...); err != nil {
		return nil, model.New500Error("sql_store.sql_user_store.get_social_account", err.Error(), nil)
	}
	return account, nil
}

func (s *sqlUserStore) GetGuestAccount(id string) (*model.GuestAccount, *model.AppError) {
	query := sqlf.
		Select("*").
		From("guest_account").
		Where("id = ?", id)

	var guestAccount model.GuestAccount
	if err := s.QueryRow(&guestAccount, query); err != nil {
		return nil, model.New500Error("sql_store.sql_user_store.get_guest_account", err.Error(), nil)
	}
	return &guestAccount, nil
}
