package rabbitmq

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

type Producer struct {
	channel      *amqp.Channel
	queueName    string
	deliveryMode uint8
	D            chan map[string]string
}

func NewProducer(connection *amqp.Connection, queueName string, deliveryMode uint8) (*Producer, error) {
	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}

	p := &Producer{channel, queueName, deliveryMode, make(chan map[string]string)}
	go p.pollAndPublish()

	return p, nil
}

func (p *Producer) pollAndPublish() error {
	for {
		message := <-p.D
		str, err := json.Marshal(message)
		if err != nil {
			continue
		}

		err = p.channel.Publish(
			DEFAULT_EXCAHNGE,
			p.queueName,
			false,
			false,
			amqp.Publishing{
				Body:         str,
				ContentType:  "application/json",
				DeliveryMode: p.deliveryMode,
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
