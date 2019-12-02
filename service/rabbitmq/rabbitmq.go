package rabbitmq

import (
	"github.com/streadway/amqp"
	"github.com/varmamsp/cello/model"
)

const (
	// Exchange names
	EXCHANGE_NAME_PHENOPOD_DIRECT = "phenopod_direct"
	EXCHANGE_NAME_PHENOPOD_DLX    = "phenopod_dlx"

	// Queue names
	QUEUE_NAME_IMPORT_PODCAST               = "import_podcast"
	QUEUE_NAME_REFRESH_PODCAST              = "refresh_podcast"
	QUEUE_NAME_CREATE_THUMBNAIL             = "create_thumbnail"
	QUEUE_NAME_SYNC_PLAYBACK                = "sync_playback"
	QUEUE_NAME_CREATE_THUMBNAIL_DEAD_LETTER = "create_thumbnail_dead_letter"

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

	// Exchanges
	if err := createExchange(EXCHANGE_NAME_PHENOPOD_DIRECT, channel); err != nil {
		return nil, err
	}

	if err := createExchange(EXCHANGE_NAME_PHENOPOD_DLX, channel); err != nil {
		return nil, err
	}

	// Queues
	if err := createAndBindQueue(
		QUEUE_NAME_IMPORT_PODCAST,
		EXCHANGE_NAME_PHENOPOD_DIRECT,
		ROUTING_KEY_IMPORT_PODCAST,
		channel,
		nil,
	); err != nil {
		return nil, err
	}

	if err := createAndBindQueue(
		QUEUE_NAME_REFRESH_PODCAST,
		EXCHANGE_NAME_PHENOPOD_DIRECT,
		ROUTING_KEY_REFRESH_PODCAST,
		channel,
		nil,
	); err != nil {
		return nil, err
	}

	if err := createAndBindQueue(
		QUEUE_NAME_CREATE_THUMBNAIL,
		EXCHANGE_NAME_PHENOPOD_DIRECT,
		ROUTING_KEY_CREATE_THUMBNAIL,
		channel,
		map[string]interface{}{
			"x-dead-letter-exchange": EXCHANGE_NAME_PHENOPOD_DLX,
		},
	); err != nil {
		return nil, err
	}

	if err := createAndBindQueue(
		QUEUE_NAME_CREATE_THUMBNAIL_DEAD_LETTER,
		EXCHANGE_NAME_PHENOPOD_DLX,
		ROUTING_KEY_CREATE_THUMBNAIL,
		channel,
		nil,
	); err != nil {
		return nil, err
	}

	if err := createAndBindQueue(
		QUEUE_NAME_SYNC_PLAYBACK,
		EXCHANGE_NAME_PHENOPOD_DIRECT,
		ROUTING_KEY_SYNC_PLAYBACK,
		channel,
		nil,
	); err != nil {
		return nil, err
	}

	return connection, nil
}

func createExchange(exchangeName string, channel *amqp.Channel) error {
	return channel.ExchangeDeclare(
		exchangeName, // name
		"direct",     // type
		true,         // durable
		false,        // delete when unused
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
}

func createAndBindQueue(queueName, exchangeName, routingKey string, channel *amqp.Channel, args map[string]interface{}) error {
	if _, err := channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		args,      // arguments
	); err != nil {
		return err
	}

	return channel.QueueBind(
		queueName,    // queue name
		routingKey,   // routing key
		exchangeName, // exchange name
		false,        // no-wait
		nil,          // arguments
	)
}

func makeAmqpUrl(config *model.Config) string {
	return "amqp://" + config.Rabbitmq.User + ":" + config.Rabbitmq.Password + "@" + config.Rabbitmq.Address + "/"
}
