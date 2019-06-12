package app

import (
	"github.com/varmamsp/cello/services/httpservice"
	"github.com/varmamsp/cello/store"
	"github.com/varmamsp/cello/store/sqlstore"
)

type App struct {
	store       store.Store
	httpservice *httpservice.Client
}

func NewApp() *App {
	return &App{
		sqlstore.NewSqlStore(),
		httpservice.NewClient(),
	}
}
