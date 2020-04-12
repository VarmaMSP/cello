package messagebroker

import (
	"github.com/streadway/amqp"
	"github.com/varmamsp/cello/model"
)

type MessageBroker interface {
	DeclareExchange(exchange, exchangeType string, durable bool) *model.AppError
	DeclareQueue(exchange, queue, routingKey, deadLetterExchange string, durable bool) *model.AppError

	NewProducer(exchange, routingKey string, deliveryMode uint8) (Producer, *model.AppError)
	NewConsumer(queue, consumer string, autoAck, exclusive bool, preFetchCount int) (Consumer, *model.AppError)
}

type Producer interface {
	Publish(o interface{}) *model.AppError
}

type Consumer interface {
	Consume(func(d amqp.Delivery))
}

func NewRabbitmqBackend(config *model.Config) (MessageBroker, *model.AppError) {
	b := &rabbitmqBackend{
		user:     config.Rabbitmq.User,
		password: config.Rabbitmq.Password,
		address:  config.Rabbitmq.Address,
	}

	if err := b.init(); err != nil {
		return nil, err
	}

	go b.reconnector()

	return b, nil
}
