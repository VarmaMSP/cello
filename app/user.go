package app

import (
	"context"
	"net/http"

	"github.com/dghubble/gologin/v2"
	facebookLogin "github.com/dghubble/gologin/v2/facebook"
	googleLogin "github.com/dghubble/gologin/v2/google"
	twitterLogin "github.com/dghubble/gologin/v2/twitter"
	"github.com/varmamsp/cello/model"
)

func (app *App) CreateUserWithGoogle(ctx context.Context) (*model.User, *model.AppError) {
	googleUser, err_ := googleLogin.UserFromContext(ctx)
	if err_ != nil {
		return nil, model.NewAppError("app.user.create_user_with_google", err_.Error(), http.StatusInternalServerError, nil)
	}

	user := &model.User{
		Name:   googleUser.Name,
		Email:  googleUser.Email,
		Gender: googleUser.Gender,
	}
	if err := app.Store.User().Save(user); err != nil {
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
	if err := app.Store.User().SaveGoogleAccount(googleAccount); err != nil {
		return nil, err
	}

	return user, nil
}

func (app *App) CreateUserWithFacebook(ctx context.Context) (*model.User, *model.AppError) {
	facebookUser, err := facebookLogin.UserFromContext(ctx)
	if err != nil {
		return nil, model.NewAppError("app.user.create_user_with_faceboook", err.Error(), http.StatusInternalServerError, nil)
	}

	user := &model.User{
		Name:  facebookUser.Name,
		Email: facebookUser.Email,
	}
	if err := app.Store.User().Save(user); err != nil {
		return nil, err
	}

	facebookAccount := &model.FacebookAccount{
		Id:     facebookUser.ID,
		UserId: user.Id,
		Name:   facebookUser.Name,
		Email:  facebookUser.Email,
	}
	if err := app.Store.User().SaveFacebookAccount(facebookAccount); err != nil {
		return nil, err
	}

	return user, nil
}

func (app *App) CreateUserWithTwitter(ctx context.Context) (*model.User, *model.AppError) {
	twitterUser, err := twitterLogin.UserFromContext(ctx)
	if err != nil {
		return nil, nil
	}

	user := &model.User{
		Name:  twitterUser.Name,
		Email: twitterUser.Email,
	}
	if err := app.Store.User().Save(user); err != nil {
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
	if err := app.Store.User().SaveTwitterAccount(twitterAccount); err != nil {
		return nil, err
	}

	return user, nil
}

func (app *App) CreateSession(ctx context.Context, user *model.User) *model.Session {
	session := &model.Session{UserId: user.Id, IsAdmin: user.IsAdmin}
	app.SessionManager.Put(ctx, "user_id", session.UserId)
	app.SessionManager.Put(ctx, "is_admin", session.IsAdmin)

	return session
}

func (app *App) GetSession(ctx context.Context) *model.Session {
	session := &model.Session{}
	session.UserId = app.SessionManager.GetString(ctx, "user_id")
	session.IsAdmin = app.SessionManager.GetInt(ctx, "is_admin")

	return session
}

func (app *App) GoogleSignIn() http.Handler {
	return googleLogin.StateHandler(gologin.DebugOnlyCookieConfig, googleLogin.LoginHandler(app.GoogleOAuthConfig, nil))
}

func (app *App) GoogleCallback(onSuccess http.Handler) http.Handler {
	return googleLogin.StateHandler(gologin.DebugOnlyCookieConfig, googleLogin.CallbackHandler(app.GoogleOAuthConfig, onSuccess, nil))
}

func (app *App) FacebookSignIn() http.Handler {
	return facebookLogin.StateHandler(gologin.DebugOnlyCookieConfig, facebookLogin.LoginHandler(app.FacebookOAuthConfig, nil))
}

func (app *App) FacebookCallback(onSuccess http.Handler) http.Handler {
	return facebookLogin.StateHandler(gologin.DebugOnlyCookieConfig, facebookLogin.CallbackHandler(app.FacebookOAuthConfig, onSuccess, nil))
}

func (app *App) TwitterSignIn() http.Handler {
	return twitterLogin.LoginHandler(app.TwitterOAuthConfig, nil)
}

func (app *App) TwitterCallback(onSuccess http.Handler) http.Handler {
	return twitterLogin.CallbackHandler(app.TwitterOAuthConfig, onSuccess, nil)
}
