package app

import "github.com/varmamsp/cello/model"

func (a *App) GetCategory(categoryId int64) (*model.Category, *model.AppError) {
	return a.Store.Category().Get(categoryId)
}

func (a *App) GetAllCategories() ([]*model.Category, *model.AppError) {
	return a.Store.Category().GetAll()
}

func (a *App) GetCategoriesByIds(categoryIds []int64) ([]*model.Category, *model.AppError) {
	return a.Store.Category().GetByIds(categoryIds)
}
