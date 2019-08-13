package main

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/varmamsp/cello/jobs"
	"github.com/varmamsp/cello/model"
	"github.com/varmamsp/cello/services/rabbitmq"
	"github.com/varmamsp/cello/store/sqlstore"
)

func main() {
	viper.SetConfigName("cello.config")
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
	conn1, err := rabbitmq.NewConnection(&config.Rabbitmq)
	if err != nil {
		fmt.Println(err)
	}
	conn2, err := rabbitmq.NewConnection(&config.Rabbitmq)
	if err != nil {
		fmt.Println(err)
	}

	jobRunner, err := jobs.NewJobRunner(store, conn1, conn2, &config.Rabbitmq.Queues)
	if err != nil {
		fmt.Println(err)
	}
	jobRunner.Start()

	var forever chan int
	<-forever
}
