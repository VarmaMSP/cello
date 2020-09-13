package app

import (
	"context"

	"github.com/varmamsp/cello/model"
	google "google.golang.org/api/oauth2/v2"
)

func (a *App) CreateUserWithGuest(guest *model.GuestAccount) (*model.User, *model.AppError) {
	if guestAccount, err := a.Store.User().GetGuestAccount(guest.Id); err == nil {
		return a.Store.User().Get(guestAccount.UserId)
	}

	user := &model.User{
		Name:         "",
		Email:        "",
		Gender:       "",
		SignInMethod: "GUEST",
	}
	if err := a.Store.User().Save(user); err != nil {
		return nil, err
	}

	guestAccount := &model.GuestAccount{
		Id:          guest.Id,
		UserId:      user.Id,
		DeviceUuid:  guest.DeviceUuid,
		DeviceOs:    guest.DeviceOs,
		DeviceModel: guest.DeviceModel,
	}
	if err := a.Store.User().SaveGuestAccount(guestAccount); err != nil {
		return nil, err
	}

	return user, nil
}

func (a *App) CreateUserWithGoogle(googleUser *google.Userinfoplus) (*model.User, *model.AppError) {
	// return user if exists
	if account, err := a.Store.User().GetGoogleAccount(googleUser.Id); err == nil {
		return a.Store.User().Get(account.UserId)
	}

	user := &model.User{
		Name:         googleUser.Name,
		Email:        googleUser.Email,
		Gender:       googleUser.Gender,
		SignInMethod: "GOOGLE",
	}
	if err := a.Store.User().Save(user); err != nil {
		return nil, err
	}

	googleAccount := &model.GoogleAccount{
		Id:         googleUser.Id,
		UserId:     user.Id,
		Email:      googleUser.Email,
		FamilyName: googleUser.FamilyName,
		Gender:     googleUser.Gender,
		GivenName:  googleUser.GivenName,
		Link:       googleUser.Link,
		Locale:     googleUser.Locale,
		Name:       googleUser.Name,
		Picture:    googleUser.Picture,
	}
	if err := a.Store.User().SaveSocialAccount("google", googleAccount); err != nil {
		return nil, err
	}

	return user, nil
}

func (a *App) LinkGoogleToUser(user *model.User, googleUser *google.Userinfoplus) *model.AppError {
	googleAccount := &model.GoogleAccount{
		Id:         googleUser.Id,
		UserId:     user.Id,
		Email:      googleUser.Email,
		FamilyName: googleUser.FamilyName,
		Gender:     googleUser.Gender,
		GivenName:  googleUser.GivenName,
		Link:       googleUser.Link,
		Locale:     googleUser.Locale,
		Name:       googleUser.Name,
		Picture:    googleUser.Picture,
	}
	if err := a.Store.User().SaveSocialAccount("google", googleAccount); err != nil {
		return err
	}

	return nil
}

func (a *App) CreateUserWithFacebook(facebookId, facebookName, facebookEmail string) (*model.User, *model.AppError) {
	// Check if user already exists
	if account, err := a.Store.User().GetFacebookAccount(facebookId); err == nil {
		return a.Store.User().Get(account.UserId)
	}

	user := &model.User{
		Name:         facebookName,
		Email:        facebookEmail,
		SignInMethod: "FACEBOOK",
	}
	if err := a.Store.User().Save(user); err != nil {
		return nil, err
	}

	facebookAccount := &model.FacebookAccount{
		Id:     facebookId,
		UserId: user.Id,
		Name:   facebookName,
		Email:  facebookEmail,
	}
	if err := a.Store.User().SaveSocialAccount("facebook", facebookAccount); err != nil {
		return nil, err
	}

	return user, nil
}

func (a *App) LinkFacebookToUser(user *model.User, facebookId, facebookName, facebookEmail string) *model.AppError {
	facebookAccount := &model.FacebookAccount{
		Id:     facebookId,
		UserId: user.Id,
		Name:   facebookName,
		Email:  facebookEmail,
	}
	if err := a.Store.User().SaveSocialAccount("facebook", facebookAccount); err != nil {
		return err
	}

	return nil
}

func (a *App) GetUser(userId int64) (*model.User, *model.AppError) {
	return a.Store.User().Get(userId)
}

func (a *App) NewSession(ctx context.Context, user *model.User) {
	session := &model.Session{UserId: user.Id, IsAdmin: user.IsAdmin}
	a.SessionManager.Put(ctx, "user_id", session.UserId)
	a.SessionManager.Put(ctx, "is_admin", session.IsAdmin)
}

func (a *App) GetSession(ctx context.Context) *model.Session {
	session := &model.Session{}

	var ok bool
	if session.UserId, ok = a.SessionManager.Get(ctx, "user_id").(int64); !ok {
		return nil
	}
	if session.IsAdmin, ok = a.SessionManager.Get(ctx, "is_admin").(int); !ok {
		return nil
	}
	return session
}

func (a *App) DeleteSession(ctx context.Context) *model.AppError {
	if err := a.SessionManager.Destroy(ctx); err != nil {
		return model.New500Error("app.delete_session", err.Error(), nil)
	}
	return nil
}
