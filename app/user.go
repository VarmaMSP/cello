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

func (app *App) GetUser(userId int64) (*model.User, *model.AppError) {
	return app.Store.User().Get(userId)
}

func (app *App) GoogleSignIn() http.Handler {
	return googleLogin.StateHandler(
		gologin.DebugOnlyCookieConfig,
		googleLogin.LoginHandler(
			app.GoogleOAuthConfig,
			nil,
		),
	)
}

func (app *App) FacebookSignIn() http.Handler {
	return facebookLogin.StateHandler(
		gologin.DebugOnlyCookieConfig,
		facebookLogin.LoginHandler(
			app.FacebookOAuthConfig,
			nil,
		),
	)
}

func (app *App) TwitterSignIn() http.Handler {
	return twitterLogin.LoginHandler(
		app.TwitterOAuthConfig,
		nil,
	)
}

func (app *App) GoogleCallback() http.Handler {
	return googleLogin.StateHandler(
		gologin.DebugOnlyCookieConfig,
		googleLogin.CallbackHandler(
			app.GoogleOAuthConfig,
			app.SessionManager.LoadAndSave(
				http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
					user, err := app.CreateUserWithGoogle(req.Context())
					if err != nil {
						w.Header().Set("Location", app.HostName)
						w.WriteHeader(http.StatusFound)
						return
					}

					app.CreateSession(req.Context(), user)
					w.Header().Set("Location", app.HostName)
					w.WriteHeader(http.StatusFound)
				}),
			),
			nil,
		),
	)
}

func (app *App) FacebookCallback() http.Handler {
	return facebookLogin.StateHandler(
		gologin.DebugOnlyCookieConfig,
		facebookLogin.CallbackHandler(
			app.FacebookOAuthConfig,
			app.SessionManager.LoadAndSave(
				http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
					user, err := app.CreateUserWithFacebook(req.Context())
					if err != nil {
						w.Header().Set("Location", app.HostName)
						w.WriteHeader(http.StatusFound)
						return
					}

					app.CreateSession(req.Context(), user)
					w.Header().Set("Location", app.HostName)
					w.WriteHeader(http.StatusFound)
				}),
			),
			nil,
		),
	)
}

func (app *App) TwitterCallback() http.Handler {
	return twitterLogin.CallbackHandler(
		app.TwitterOAuthConfig,
		app.SessionManager.LoadAndSave(
			http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				user, err := app.CreateUserWithTwitter(req.Context())
				if err != nil {
					w.Header().Set("Location", app.HostName)
					w.WriteHeader(http.StatusFound)
					return
				}

				app.CreateSession(req.Context(), user)
				w.Header().Set("Location", app.HostName)
				w.WriteHeader(http.StatusFound)
			}),
		),
		nil,
	)
}

func (app *App) CreateUserWithGoogle(ctx context.Context) (*model.User, *model.AppError) {
	googleUser, err := googleLogin.UserFromContext(ctx)
	if err != nil {
		return nil, model.NewAppError("app.user.create_user_with_google", err.Error(), http.StatusInternalServerError, nil)
	}

	// Check if user already exists
	if account, err := app.Store.User().GetSocialAccount("google", googleUser.Id); err == nil {
		googleAccount, _ := account.(*model.GoogleAccount)
		return app.Store.User().Get(googleAccount.UserId)
	}

	user := &model.User{
		Name:         googleUser.Name,
		Email:        googleUser.Email,
		Gender:       googleUser.Gender,
		SignInMethod: "GOOGLE",
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
	if err := app.Store.User().SaveSocialAccount("google", googleAccount); err != nil {
		return nil, err
	}

	return user, nil
}

func (app *App) CreateUserWithFacebook(ctx context.Context) (*model.User, *model.AppError) {
	facebookUser, err := facebookLogin.UserFromContext(ctx)
	if err != nil {
		return nil, model.NewAppError("app.user.create_user_with_faceboook", err.Error(), http.StatusInternalServerError, nil)
	}

	// Check if user already exists
	if account, err := app.Store.User().GetSocialAccount("facebook", facebookUser.ID); err == nil {
		facebookAccount, _ := account.(*model.FacebookAccount)
		return app.Store.User().Get(facebookAccount.UserId)
	}

	user := &model.User{
		Name:         facebookUser.Name,
		Email:        facebookUser.Email,
		SignInMethod: "FACEBOOK",
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
	if err := app.Store.User().SaveSocialAccount("facebook", facebookAccount); err != nil {
		return nil, err
	}

	return user, nil
}

func (app *App) CreateUserWithTwitter(ctx context.Context) (*model.User, *model.AppError) {
	twitterUser, err := twitterLogin.UserFromContext(ctx)
	if err != nil {
		return nil, nil
	}

	// Check if user already exists
	if account, err := app.Store.User().GetSocialAccount("twitter", twitterUser.IDStr); err == nil {
		twitterAccount, _ := account.(*model.TwitterAccount)
		return app.Store.User().Get(twitterAccount.UserId)
	}

	user := &model.User{
		Name:         twitterUser.Name,
		Email:        twitterUser.Email,
		SignInMethod: "TWITTER",
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
	if err := app.Store.User().SaveSocialAccount("twitter", twitterAccount); err != nil {
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

	var ok bool
	if session.UserId, ok = app.SessionManager.Get(ctx, "user_id").(int64); !ok {
		return nil
	}
	if session.IsAdmin, ok = app.SessionManager.Get(ctx, "is_admin").(int); !ok {
		return nil
	}
	return session
}

func (app *App) DeleteSession(ctx context.Context) *model.AppError {
	err := app.SessionManager.Destroy(ctx)
	if err != nil {
		return model.NewAppError("app.delete_session", err.Error(), http.StatusInternalServerError, nil)
	}
	return nil
}
