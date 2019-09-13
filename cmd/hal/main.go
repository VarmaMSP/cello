package main

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/varmamsp/cello/api"
	"github.com/varmamsp/cello/model"
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

	api := api.NewApi(config)
	api.ListenAndServe()
}
