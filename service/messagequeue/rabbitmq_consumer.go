package messagequeue

import (
	"github.com/streadway/amqp"
)

type consumerSupplier struct {
	connection    *amqp.Connection
	channel       *amqp.Channel
	queue         string
	consumer      string
	autoAck       bool
	exclusive     bool
	preFetchCount int
	deliveries    <-chan amqp.Delivery
}

func (c *consumerSupplier) init() error {
	if channel, err := c.connection.Channel(); err != nil {
		return err
	} else if err := channel.Qos(c.preFetchCount, 0, false); err != nil {
		return err
	} else {
		c.channel = channel
	}

	if d, err := c.channel.Consume(c.queue, c.consumer, c.autoAck, c.exclusive, false, false, nil); err != nil {
		return err
	} else {
		c.deliveries = d
	}

	return nil
}

func (c *consumerSupplier) Consume() <-chan amqp.Delivery {
	return c.deliveries
}
