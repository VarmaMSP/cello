package message_queue

import (
	"github.com/streadway/amqp"
)

type Broker interface {
	NewProducer(exchange, routingKey string, deliveryMode uint8) (Producer, error)
	NewConsumer(queue, consumer string, autoAck, exclusive bool, preFetchCount int) (Consumer, error)
}

type Producer interface {
	Publish(o interface{}) error
}

type Consumer interface {
	Consume() <-chan amqp.Delivery
}
