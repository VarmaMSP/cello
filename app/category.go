package app

import (
	"github.com/varmamsp/cello/model"
)

func (app *App) GetCategoriesByIds(categoryIds []int64) ([]*model.Category, *model.AppError) {
	return app.Store.Category().GetByIds(categoryIds)
}
