package rabbitmq

import (
	"github.com/streadway/amqp"
)

type Consumer struct {
	channel   *amqp.Channel
	queueName string
	D         <-chan amqp.Delivery
}

func NewConsumer(connection *amqp.Connection, queueName string) (*Consumer, error) {
	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}

	d, err := channel.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		return nil, err
	}

	return &Consumer{channel, queueName, d}, nil
}

func (c *Consumer) Close() {
	c.channel.Close()
}
