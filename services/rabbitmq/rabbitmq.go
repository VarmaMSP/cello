package rabbitmq

import "github.com/streadway/amqp"

const (
	DEFAULT_EXCAHNGE = ""

	QUEUE_NEW_PODCAST  = "new_podcast"
	QUEUE_RESIZE_IMAGE = "resize_image"

	Transient  = 1
	Persistent = 2
)

func NewConnection() (*amqp.Connection, error) {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}

	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}
	defer channel.Close()

	if err := createQueue(QUEUE_NEW_PODCAST, channel); err != nil {
		return nil, err
	}
	if err := createQueue(QUEUE_RESIZE_IMAGE, channel); err != nil {
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
