package messagebroker

import (
	"net/http"
	"time"

	"github.com/streadway/amqp"
	"github.com/varmamsp/cello/model"
)

type rabbitmqBackend struct {
	user     string
	password string
	address  string

	producerConn *amqp.Connection
	consumerConn *amqp.Connection

	producers []*rabbitmqProducer
	consumers []*rabbitmqConsumer
}

func (b *rabbitmqBackend) init() *model.AppError {
	amqpUrl := "amqp://" + b.user + ":" + b.password + "@" + b.address + "/"

	if connection, err := amqp.Dial(amqpUrl); err != nil {
		return model.NewAppError("rabbimq_backend.init", err.Error(), http.StatusInternalServerError, nil)
	} else {
		b.producerConn = connection
	}

	if connection, err := amqp.Dial(amqpUrl); err != nil {
		return model.NewAppError("rabbimq_backend.init", err.Error(), http.StatusInternalServerError, nil)
	} else {
		b.consumerConn = connection
	}

	return nil
}

func (b *rabbitmqBackend) reconnector() {
	for {
		select {
		case _, ok := <-b.producerConn.NotifyClose(make(chan *amqp.Error)):
			if !ok {
				return
			}
			break

		case _, ok := <-b.consumerConn.NotifyClose(make(chan *amqp.Error)):
			if !ok {
				return
			}
			break
		}

		for {
			time.Sleep(40 * time.Second)
			if err := b.init(); err != nil {
				continue
			}

			// recover
			for _, producer := range b.producers {
				producer.recover <- struct{}{}
			}
			for _, consumer := range b.consumers {
				consumer.recover <- struct{}{}
			}
		}
	}
}

func (b *rabbitmqBackend) getProducerConn() *amqp.Connection {
	return b.producerConn
}

func (b *rabbitmqBackend) getConsumerConn() *amqp.Connection {
	return b.consumerConn
}

func (b *rabbitmqBackend) DeclareExchange(exchange, exchangeType string, durable bool) *model.AppError {
	channel, err := b.producerConn.Channel()
	if err != nil {
		return model.NewAppError("rabbitmq_backend.declare_exchange", err.Error(), http.StatusInternalServerError, nil)
	}
	defer channel.Close()

	if err := channel.ExchangeDeclare(exchange, exchangeType, durable, false, false, false, nil); err != nil {
		return model.NewAppError("rabbitmq_backend.declare_exchange", err.Error(), http.StatusInternalServerError, nil)
	}

	return nil
}

func (b *rabbitmqBackend) DeclareQueue(exchange, queue, routingKey, deadLetterExchange string, durable bool) *model.AppError {
	channel, err := b.producerConn.Channel()
	if err != nil {
		return model.NewAppError("rabbitmq_backend.declare_queue", err.Error(), http.StatusInternalServerError, nil)
	}
	defer channel.Close()

	var args map[string]interface{}
	if deadLetterExchange != "" {
		args = map[string]interface{}{"x-dead-letter-exchange": deadLetterExchange}
	}

	if _, err := channel.QueueDeclare(queue, durable, false, false, false, args); err != nil {
		return model.NewAppError("rabbitmq_backend.declare_queue", err.Error(), http.StatusInternalServerError, nil)
	}

	if err := channel.QueueBind(queue, routingKey, exchange, false, nil); err != nil {
		return model.NewAppError("rabbitmq_backend.declare_queue", err.Error(), http.StatusInternalServerError, nil)
	}

	return nil
}

func (b *rabbitmqBackend) NewProducer(exchange, routingKey string, deliveryMode uint8) (Producer, *model.AppError) {
	p := &rabbitmqProducer{
		getConnection: b.getProducerConn,
		recover:       make(chan struct{}),
		exchange:      exchange,
		routingKey:    routingKey,
		deliveryMode:  deliveryMode,
	}

	if err := p.init(); err != nil {
		return nil, err
	} else {
		b.producers = append(b.producers, p)
		go p.reconnector()
		return p, nil
	}
}

func (b *rabbitmqBackend) NewConsumer(queue, consumer string, autoAck, exclusive bool, preFetchCount int) (Consumer, *model.AppError) {
	c := &rabbitmqConsumer{
		getConnection: b.getConsumerConn,
		recover:       make(chan struct{}),
		queue:         queue,
		consumer:      consumer,
		autoAck:       autoAck,
		exclusive:     exclusive,
		preFetchCount: preFetchCount,
	}

	if err := c.init(); err != nil {
		return nil, err
	} else {
		b.consumers = append(b.consumers, c)
		go c.reconnector()
		return c, nil
	}
}
