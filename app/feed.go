package app

import "github.com/varmamsp/cello/model"

func (app *App) GetFeed(id string) (*model.Feed, *model.AppError) {
	return app.Store.Feed().Get(id)
}
