package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/service/cache"
	"github.com/varmamsp/cello/service/filestorage"
	"github.com/varmamsp/cello/service/messagequeue"
	"github.com/varmamsp/cello/service/searchengine"
	"github.com/varmamsp/cello/service/sqldb"
	"github.com/varmamsp/cello/store"
	"github.com/varmamsp/cello/store/searchlayer"
	"github.com/varmamsp/cello/store/sqlstore"
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

	// Config
	var config model.Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(err.Error())
		return
	}

	// Store
	var store store.Store
	if db, err := sqldb.NewBroker(&config); err != nil {
		fmt.Println(err.Error())
		return
	} else if s, err := sqlstore.NewSqlStore(db); err != nil {
		fmt.Println(err.Error())
		return
	} else {
		store = s
	}

	// Searchengine
	var seBroker searchengine.Broker
	if se, err := searchengine.NewBroker(&config); err != nil {
		fmt.Println(err.Error())
		return
	} else {
		store = searchlayer.NewSearchLayer(store, se)
		seBroker = se
	}

	// Messagequeue
	var mqBroker messagequeue.Broker
	if mq, err := messagequeue.NewBroker(&config); err != nil {
		fmt.Println(err.Error())
		return
	} else {
		mqBroker = mq
	}

	// Filestorage
	var fsBroker filestorage.Broker
	if fs, err := filestorage.NewBroker(&config); err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fsBroker = fs
	}

	// Cache
	var cacheBroker cache.Broker
	if cache, err := cache.NewBroker(&config); err != nil {
		fmt.Println(err.Error())
		return
	} else {
		cacheBroker = cache
	}

	// Logger
	var logger zerolog.Logger
	if config.Env == "dev" {
		logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()
	} else {
		logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	}
}
