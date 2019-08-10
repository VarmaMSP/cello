package main

import (
	"fmt"

	"github.com/varmamsp/cello/jobs"
	"github.com/varmamsp/cello/services/rabbitmq"
	"github.com/varmamsp/cello/store/sqlstore"
)

func main() {
	store := sqlstore.NewSqlStore()
	conn1, err := rabbitmq.NewConnection()
	if err != nil {
		fmt.Println(err)
	}

	conn2, err := rabbitmq.NewConnection()
	if err != nil {
		fmt.Println(err)
	}

	jobRunner, err := jobs.NewJobRunner(store, conn1, conn2)
	if err != nil {
		fmt.Println(err)
	}

	jobRunner.Start()

	var forever chan int
	<-forever
}
