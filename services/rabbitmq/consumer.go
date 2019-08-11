package rabbitmq

import (
	"github.com/streadway/amqp"
)

type Consumer struct {
	channel *amqp.Channel
	D       <-chan amqp.Delivery
}

type ConsumerOpts struct {
	QueueName     string
	ConsumerName  string
	AutoAck       bool
	PreFetchCount int
}

func NewConsumer(conn *amqp.Connection, opts *ConsumerOpts) (*Consumer, error) {
	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	if err := channel.Qos(opts.PreFetchCount, 0, false); err != nil {
		return nil, err
	}

	d, err := channel.Consume(
		opts.QueueName,    // queue
		opts.ConsumerName, // consumer
		opts.AutoAck,      // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)
	if err != nil {
		return nil, err
	}

	return &Consumer{channel, d}, nil
}

func (c *Consumer) Close() {
	c.channel.Close()
}
