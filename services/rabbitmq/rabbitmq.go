package rabbitmq

import (
	"github.com/streadway/amqp"
	"github.com/varmamsp/cello/model"
)

const (
	DefaultExchange = ""
)

func NewConnection(config *model.RabbitmqConfig) (*amqp.Connection, error) {
	connection, err := amqp.Dial(makeAmqpUrl(config))
	if err != nil {
		return nil, err
	}

	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}
	defer channel.Close()

	if err := createQueue(model.QUEUE_NAME_IMPORT_PODCAST, channel); err != nil {
		return nil, err
	}
	if err := createQueue(model.QUEUE_NAME_SCHEDULED_JOB_CALL, channel); err != nil {
		return nil, err
	}
	if err := createQueue(model.QUEUE_NAME_REFRESH_PODCAST, channel); err != nil {
		return nil, err
	}
	return connection, nil
}

func createQueue(name string, channel *amqp.Channel) error {
	_, err := channel.QueueDeclare(
		name,  // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	return err
}

func makeAmqpUrl(config *model.RabbitmqConfig) string {
	return "amqp://" + config.User + ":" + config.Password + "@" + config.Address + "/"
}
