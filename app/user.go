package app

import (
	"context"
	"net/http"

	"github.com/dghubble/gologin/v2"
	"github.com/dghubble/gologin/v2/google"
	"github.com/varmamsp/cello/model"
	googleOAuth "google.golang.org/api/oauth2/v2"
)

func (app *App) CreateUserWithGoogle(googleUser *googleOAuth.Userinfoplus) (*model.User, *model.AppError) {
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
	return google.StateHandler(gologin.DebugOnlyCookieConfig, google.LoginHandler(app.GoogleOAuthConfig, nil))
}

func (app *App) GoogleCallback(onSuccess http.Handler) http.Handler {
	return google.StateHandler(gologin.DebugOnlyCookieConfig, google.CallbackHandler(app.GoogleOAuthConfig, onSuccess, nil))
}
