package app

import (
	"github.com/varmamsp/cello/store"
)

type App struct {
	store store.Store
}

func NewApp(store store.Store) *App {
	return &App{
		store,
	}
}
