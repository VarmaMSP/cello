package messagebroker

import (
	"encoding/json"
	"net/http"

	"github.com/streadway/amqp"
	"github.com/varmamsp/cello/model"
)

type rabbitmqProducer struct {
	getConnection func() *amqp.Connection
	recover       chan struct{}

	channel      *amqp.Channel
	exchange     string
	routingKey   string
	deliveryMode uint8
}

func (p *rabbitmqProducer) init() *model.AppError {
	connection := p.getConnection()
	if connection == nil {
		return model.NewAppError("rabbitmq_producer.init", "no connection", http.StatusInternalServerError, nil)
	}

	if channel, err := connection.Channel(); err != nil {
		return model.NewAppError("rabbitmq_producer.init", err.Error(), http.StatusInternalServerError, nil)
	} else {
		p.channel = channel
	}

	return nil
}

func (p *rabbitmqProducer) reconnector() {
	for {
		<-p.recover
		p.init()
	}
}

func (p *rabbitmqProducer) Publish(o interface{}) *model.AppError {
	str, err := json.Marshal(o)
	if err != nil {
		return model.NewAppError("rabbitmq.producer.publish", err.Error(), http.StatusBadRequest, nil)
	}

	if err := p.channel.Publish(
		p.exchange,
		p.routingKey,
		false, // immediate
		false, // publishing
		amqp.Publishing{
			Body:         str,
			ContentType:  "application/json",
			DeliveryMode: p.deliveryMode,
		},
	); err != nil {
		return model.NewAppError("rabbitmq.producer.publish", err.Error(), http.StatusInternalServerError, nil)
	}

	return nil
}
