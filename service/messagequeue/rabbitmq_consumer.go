package messagequeue

import (
	"errors"

	"github.com/streadway/amqp"
)

type consumerSupplier struct {
	getConnection func() *amqp.Connection

	channel       *amqp.Channel
	queue         string
	consumer      string
	autoAck       bool
	exclusive     bool
	preFetchCount int

	deliveries       <-chan amqp.Delivery
	handleDeliveryF  func(d amqp.Delivery)
	stopConsumptionC chan struct{}
}

func (c *consumerSupplier) init() error {
	connection := c.getConnection()
	if connection == nil {
		return errors.New("No Connection provided")
	}

	if channel, err := connection.Channel(); err != nil {
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

func (c *consumerSupplier) startConsumption() {
	if c.handleDeliveryF == nil {
		return
	}

	c.stopConsumptionC = make(chan struct{}, 0)
	for {
		select {
		case d := <-c.deliveries:
			c.handleDeliveryF(d)

		case <-c.stopConsumptionC:
			c.stopConsumptionC = nil
			return
		}
	}
}

func (c *consumerSupplier) stopConsumption() {
	if c.stopConsumptionC != nil {
		c.stopConsumptionC <- struct{}{}
	}
}

func (c *consumerSupplier) recover() {
	c.stopConsumption()
	c.init()
	c.startConsumption()
}

func (c *consumerSupplier) Consume(f func(d amqp.Delivery)) {
	c.stopConsumption()
	c.handleDeliveryF = f
	go c.startConsumption()
}
