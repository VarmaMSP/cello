package app

import (
	"context"
	"net/http"

	facebookLogin "github.com/dghubble/gologin/v2/facebook"
	googleLogin "github.com/dghubble/gologin/v2/google"
	twitterLogin "github.com/dghubble/gologin/v2/twitter"
	"github.com/varmamsp/cello/model"
)

func (a *App) CreateUserWithGoogle(ctx context.Context) (*model.User, *model.AppError) {
	googleUser, err := googleLogin.UserFromContext(ctx)
	if err != nil {
		return nil, model.NewAppError("app.user.create_user_with_google", err.Error(), http.StatusInternalServerError, nil)
	}

	// return user if exists
	if account, err := a.Store.User().GetSocialAccount("google", googleUser.Id); err == nil {
		googleAccount, _ := account.(*model.GoogleAccount)
		return a.Store.User().Get(googleAccount.UserId)
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

func (a *App) CreateUserWithFacebook(ctx context.Context) (*model.User, *model.AppError) {
	facebookUser, err := facebookLogin.UserFromContext(ctx)
	if err != nil {
		return nil, model.NewAppError("app.user.create_user_with_faceboook", err.Error(), http.StatusInternalServerError, nil)
	}

	// Check if user already exists
	if account, err := a.Store.User().GetSocialAccount("facebook", facebookUser.ID); err == nil {
		facebookAccount, _ := account.(*model.FacebookAccount)
		return a.Store.User().Get(facebookAccount.UserId)
	}

	user := &model.User{
		Name:         facebookUser.Name,
		Email:        facebookUser.Email,
		SignInMethod: "FACEBOOK",
	}
	if err := a.Store.User().Save(user); err != nil {
		return nil, err
	}

	facebookAccount := &model.FacebookAccount{
		Id:     facebookUser.ID,
		UserId: user.Id,
		Name:   facebookUser.Name,
		Email:  facebookUser.Email,
	}
	if err := a.Store.User().SaveSocialAccount("facebook", facebookAccount); err != nil {
		return nil, err
	}

	return user, nil
}

func (a *App) CreateUserWithTwitter(ctx context.Context) (*model.User, *model.AppError) {
	twitterUser, err := twitterLogin.UserFromContext(ctx)
	if err != nil {
		return nil, nil
	}

	// Check if user already exists
	if account, err := a.Store.User().GetSocialAccount("twitter", twitterUser.IDStr); err == nil {
		twitterAccount, _ := account.(*model.TwitterAccount)
		return a.Store.User().Get(twitterAccount.UserId)
	}

	user := &model.User{
		Name:         twitterUser.Name,
		Email:        twitterUser.Email,
		SignInMethod: "TWITTER",
	}
	if err := a.Store.User().Save(user); err != nil {
		return nil, err
	}

	twitterAccount := &model.TwitterAccount{
		Id:             twitterUser.IDStr,
		UserId:         user.Id,
		Name:           twitterUser.Name,
		ScreenName:     twitterUser.ScreenName,
		Location:       twitterUser.Location,
		Url:            twitterUser.URL,
		Description:    twitterUser.Description,
		FollowersCount: twitterUser.FollowersCount,
		FriendsCount:   twitterUser.FriendsCount,
		ProfileImage:   twitterUser.ProfileImageURLHttps,
	}
	if twitterUser.Verified {
		twitterAccount.Verified = 1
	}
	if err := a.Store.User().SaveSocialAccount("twitter", twitterAccount); err != nil {
		return nil, err
	}

	return user, nil
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
