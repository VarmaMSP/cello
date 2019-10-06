package main

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/varmamsp/cello/api"
	"github.com/varmamsp/cello/model"
)

const (
	CONFIG_NAME = "api-server.conf"
	BUILD_DIR   = "/usr/local/api-server"
)

func main() {
	viper.SetConfigName(CONFIG_NAME)
	viper.AddConfigPath(BUILD_DIR)
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../../config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err.Error())
		return
	}

	var config model.Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(err.Error())
		return
	}

	api, err := api.NewApi(config)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	api.ListenAndServe()
}
