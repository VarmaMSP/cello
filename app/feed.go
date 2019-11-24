package app

import "github.com/varmamsp/cello/model"

func (app *App) GetFeed(id int64) (*model.Feed, *model.AppError) {
	return app.Store.Feed().Get(id)
}
