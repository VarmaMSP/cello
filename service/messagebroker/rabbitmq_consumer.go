package messagebroker

import (
	"net/http"

	"github.com/streadway/amqp"
	"github.com/varmamsp/cello/model"
)

type rabbitmqConsumer struct {
	getConnection func() *amqp.Connection
	recover       chan struct{}

	channel       *amqp.Channel
	queue         string
	consumer      string
	autoAck       bool
	exclusive     bool
	preFetchCount int

	deliveries       <-chan amqp.Delivery
	handleDelivery   func(d amqp.Delivery)
	stopConsumptionC chan struct{}
}

func (c *rabbitmqConsumer) init() *model.AppError {
	connection := c.getConnection()
	if connection == nil {
		return model.NewAppError("rabbitmq_consumer.init", "no connection", http.StatusInternalServerError, nil)
	}

	if channel, err := connection.Channel(); err != nil {
		return model.NewAppError("rabbitmq_consumer.init", err.Error(), http.StatusInternalServerError, nil)
	} else if err := channel.Qos(c.preFetchCount, 0, false); err != nil {
		return model.NewAppError("rabbitmq_consumer.init", err.Error(), http.StatusInternalServerError, nil)
	} else {
		c.channel = channel
	}

	if d, err := c.channel.Consume(c.queue, c.consumer, c.autoAck, c.exclusive, false, false, nil); err != nil {
		return model.NewAppError("rabbitmq_consumer.init", err.Error(), http.StatusInternalServerError, nil)
	} else {
		c.deliveries = d
	}

	return nil
}

func (c *rabbitmqConsumer) reconnector() {
	for {
		<-c.recover

		c.stopConsumption()
		c.init()
		go c.startConsumption()
	}
}

func (c *rabbitmqConsumer) startConsumption() {
	if c.handleDelivery == nil {
		return
	}

	c.stopConsumptionC = make(chan struct{}, 0)
	for {
		select {
		case d := <-c.deliveries:
			c.handleDelivery(d)

		case <-c.stopConsumptionC:
			c.stopConsumptionC = nil
			return
		}
	}
}

func (c *rabbitmqConsumer) stopConsumption() {
	if c.stopConsumptionC != nil {
		c.stopConsumptionC <- struct{}{}
	}
}

func (c *rabbitmqConsumer) Consume(f func(d amqp.Delivery)) {
	c.stopConsumption()
	c.handleDelivery = f
	go c.startConsumption()
}
