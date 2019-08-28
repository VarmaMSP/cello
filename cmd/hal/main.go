package main

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/varmamsp/cello/api"
	"github.com/varmamsp/cello/app"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/store/sqlstore"
)

func main() {
	viper.SetConfigName("cello.conf")
	viper.AddConfigPath("./configs")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err.Error())
		return
	}

	var config model.Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(err.Error())
		return
	}

	store := sqlstore.NewSqlStore(&config.Mysql)
	app := app.NewApp(store)
	api := api.NewApi(app)
	api.ListenAndServe()
}
