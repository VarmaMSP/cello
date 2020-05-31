package messagequeue

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

type producerSupplier struct {
	connection   *amqp.Connection
	channel      *amqp.Channel
	exchange     string
	routingKey   string
	deliveryMode uint8
}

func (p *producerSupplier) init() error {
	if channel, err := p.connection.Channel(); err != nil {
		return err
	} else {
		p.channel = channel
	}
	return nil
}

func (p *producerSupplier) Publish(o interface{}) error {
	str, err := json.Marshal(o)
	if err != nil {
		return err
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
		return err
	}

	return nil
}
