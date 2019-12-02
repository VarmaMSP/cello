package rabbitmq

import (
	"github.com/streadway/amqp"
	"github.com/varmamsp/cello/model"
)

const (
	// Exchange names
	EXCHANGE_NAME_PHENOPOD_DIRECT = "phenopod_direct"

	// Queue names
	QUEUE_NAME_IMPORT_PODCAST   = "import_podcast"
	QUEUE_NAME_REFRESH_PODCAST  = "refresh_podcast"
	QUEUE_NAME_CREATE_THUMBNAIL = "create_thumbnail"
	QUEUE_NAME_SYNC_PLAYBACK    = "sync_playback"

	// Routing keys
	ROUTING_KEY_IMPORT_PODCAST   = "rk_import_podcast"
	ROUTING_KEY_REFRESH_PODCAST  = "rk_refresh_podcast"
	ROUTING_KEY_CREATE_THUMBNAIL = "rk_create_thumbnail"
	ROUTING_KEY_SYNC_PLAYBACK    = "rk_sync_playback"
)

func NewConnection(config *model.Config) (*amqp.Connection, error) {
	connection, err := amqp.Dial(makeAmqpUrl(config))
	if err != nil {
		return nil, err
	}

	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}
	defer channel.Close()

	// Phenopod direct
	if err := createDirectExchange(channel); err != nil {
		return nil, err
	}
	if err := createAndBindQueue(QUEUE_NAME_IMPORT_PODCAST, ROUTING_KEY_CREATE_THUMBNAIL, channel); err != nil {
		return nil, err
	}
	if err := createAndBindQueue(QUEUE_NAME_REFRESH_PODCAST, ROUTING_KEY_REFRESH_PODCAST, channel); err != nil {
		return nil, err
	}
	if err := createAndBindQueue(QUEUE_NAME_CREATE_THUMBNAIL, ROUTING_KEY_CREATE_THUMBNAIL, channel); err != nil {
		return nil, err
	}
	if err := createAndBindQueue(QUEUE_NAME_SYNC_PLAYBACK, ROUTING_KEY_SYNC_PLAYBACK, channel); err != nil {
		return nil, err
	}
	return connection, nil
}

func createDirectExchange(channel *amqp.Channel) error {
	return channel.ExchangeDeclare(
		EXCHANGE_NAME_PHENOPOD_DIRECT, // name
		"direct",                      // type
		true,                          // durable
		false,                         // delete when unused
		false,                         // internal
		false,                         // no-wait
		nil,                           // arguments
	)
}

func createAndBindQueue(queueName string, routingKey string, channel *amqp.Channel) error {
	if _, err := channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	); err != nil {
		return err
	}

	return channel.QueueBind(
		queueName,                     // queue name
		routingKey,                    // routing key
		EXCHANGE_NAME_PHENOPOD_DIRECT, // exchange name
		false,                         // no-wait
		nil,                           // arguments
	)
}

func makeAmqpUrl(config *model.Config) string {
	return "amqp://" + config.Rabbitmq.User + ":" + config.Rabbitmq.Password + "@" + config.Rabbitmq.Address + "/"
}
