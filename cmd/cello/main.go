package main

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/server"
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

	svr, err := server.New(&config)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	svr.Start()
}
