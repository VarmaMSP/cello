package rabbitmq

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

type Producer struct {
	channel *amqp.Channel
	D       chan interface{}
}

type ProducerOpts struct {
	ExchangeName string
	QueueName    string
	DeliveryMode uint8
}

func NewProducer(connection *amqp.Connection, opts *ProducerOpts) (*Producer, error) {
	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}

	p := &Producer{channel, make(chan interface{})}
	go p.pollAndPublish(opts)

	return p, nil
}

func (p *Producer) pollAndPublish(opts *ProducerOpts) error {
	for {
		message := <-p.D
		str, err := json.Marshal(message)
		if err != nil {
			continue
		}

		err = p.channel.Publish(
			opts.ExchangeName, // exchange
			opts.QueueName,    // exchange key
			false,             // immediate
			false,             // publishing
			amqp.Publishing{
				Body:         str,
				ContentType:  "application/json",
				DeliveryMode: opts.DeliveryMode,
			},
		)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func (p *Producer) Close() {
	p.channel.Close()
}
