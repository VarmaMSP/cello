package main

import (
	"fmt"

	"github.com/streadway/amqp"
	"github.com/varmamsp/cello/jobs/itunescrawler"
	"github.com/varmamsp/cello/jobs/podcastimport"
	"github.com/varmamsp/cello/services/rabbitmq"
	"github.com/varmamsp/cello/store/sqlstore"
)

func main() {
	store := sqlstore.NewSqlStore()
	conn, err := rabbitmq.NewConnection()
	if err != nil {
		fmt.Println(err)
	}

	producer, err := rabbitmq.NewProducer(conn, rabbitmq.QUEUE_NEW_PODCAST, amqp.Transient)
	if err != nil {
		fmt.Println(err)
	}
	p, err := itunescrawler.New(store, producer, 10)
	if err != nil {
		fmt.Println(err)
	}
	p.Run()

	consumer, err := rabbitmq.NewConsumer(conn, rabbitmq.QUEUE_NEW_PODCAST)
	if err != nil {
		fmt.Println(err)
	}
	c, err := podcastimport.New(store, consumer, 10)
	if err != nil {
		fmt.Println(err)
	}
	go c.Run()

	var forever chan int
	<-forever
}
