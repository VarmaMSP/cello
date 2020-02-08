package app

import (
	"github.com/varmamsp/cello/model"
)

func (app *App) GetCategory(categoryId int64) (*model.Category, *model.AppError) {
	return app.Store.Category().Get(categoryId)
}

func (app *App) GetAllCategories() ([]*model.Category, *model.AppError) {
	return app.Store.Category().GetAll()
}

func (app *App) GetCategoriesByIds(categoryIds []int64) ([]*model.Category, *model.AppError) {
	return app.Store.Category().GetByIds(categoryIds)
}
