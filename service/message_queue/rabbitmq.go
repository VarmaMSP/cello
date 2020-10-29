package message_queue

import (
	"fmt"

	"github.com/streadway/amqp"
	"github.com/varmamsp/cello/model"
)

type supplier struct {
	user     string
	password string
	address  string

	producerConn *amqp.Connection
	consumerConn *amqp.Connection
}

func NewBroker(config *model.Config) (Broker, error) {
	splr := &supplier{
		user:     config.Rabbitmq.User,
		password: config.Rabbitmq.Password,
		address:  config.Rabbitmq.Address,
	}

	// Init
	if err := splr.init(); err != nil {
		return nil, err
	}

	// Exchanges
	if err := splr.createExchange(EXCHANGE_PHENOPOD_DIRECT); err != nil {
		return nil, err
	}

	if err := splr.createExchange(EXCHANGE_PHENOPOD_DLX); err != nil {
		return nil, err
	}

	// Queues
	if err := splr.createAndBindQueue(
		QUEUE_IMPORT_PODCAST,
		EXCHANGE_PHENOPOD_DIRECT,
		ROUTING_KEY_IMPORT_PODCAST,
		nil,
	); err != nil {
		return nil, err
	}

	if err := splr.createAndBindQueue(
		QUEUE_REFRESH_PODCAST,
		EXCHANGE_PHENOPOD_DIRECT,
		ROUTING_KEY_REFRESH_PODCAST,
		nil,
	); err != nil {
		return nil, err
	}

	if err := splr.createAndBindQueue(
		QUEUE_CREATE_THUMBNAIL,
		EXCHANGE_PHENOPOD_DIRECT,
		ROUTING_KEY_CREATE_THUMBNAIL,
		map[string]interface{}{
			"x-dead-letter-exchange": EXCHANGE_PHENOPOD_DLX,
		},
	); err != nil {
		return nil, err
	}

	if err := splr.createAndBindQueue(
		QUEUE_CREATE_THUMBNAIL_DEAD_LETTER,
		EXCHANGE_PHENOPOD_DLX,
		ROUTING_KEY_CREATE_THUMBNAIL,
		nil,
	); err != nil {
		return nil, err
	}

	if err := splr.createAndBindQueue(
		QUEUE_SYNC_PLAYBACK,
		EXCHANGE_PHENOPOD_DIRECT,
		ROUTING_KEY_SYNC_PLAYBACK,
		nil,
	); err != nil {
		return nil, err
	}

	if err := splr.createAndBindQueue(
		QUEUE_SCHEDULED_TASK,
		EXCHANGE_PHENOPOD_DIRECT,
		ROUTING_KEY_SCHEDULED_TASK,
		nil,
	); err != nil {
		return nil, err
	}

	return splr, nil
}

func (splr *supplier) init() error {
	amqpUrl := fmt.Sprintf("amqp://%s:%s@%s/", splr.user, splr.password, splr.address)

	if connection, err := amqp.Dial(amqpUrl); err != nil {
		return err
	} else {
		splr.producerConn = connection
	}

	if connection, err := amqp.Dial(amqpUrl); err != nil {
		return err
	} else {
		splr.consumerConn = connection
	}

	return nil
}

func (splr *supplier) createExchange(exchange string) error {
	channel, err := splr.producerConn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	return channel.ExchangeDeclare(
		exchange, // name
		"direct", // type
		true,     // durable
		false,    // delete when unused
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
}

func (splr *supplier) createAndBindQueue(queue, exchange, routingKey string, args map[string]interface{}) error {
	channel, err := splr.producerConn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	if _, err := channel.QueueDeclare(
		queue, // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		args,  // arguments
	); err != nil {
		return err
	}

	return channel.QueueBind(
		queue,      // queue name
		routingKey, // routing key
		exchange,   // exchange name
		false,      // no-wait
		nil,        // arguments
	)
}

func (splr *supplier) getProducerConn() *amqp.Connection {
	return splr.producerConn
}

func (splr *supplier) getConsumerConn() *amqp.Connection {
	return splr.consumerConn
}

func (splr *supplier) NewProducer(exchange, routingKey string, deliveryMode uint8) (Producer, error) {
	p := &producerSupplier{
		connection:   splr.producerConn,
		exchange:     exchange,
		routingKey:   routingKey,
		deliveryMode: deliveryMode,
	}

	if err := p.init(); err != nil {
		return nil, err
	}

	return p, nil
}

func (splr *supplier) NewConsumer(queue, consumer string, autoAck, exclusive bool, preFetchCount int) (Consumer, error) {
	c := &consumerSupplier{
		connection:    splr.consumerConn,
		queue:         queue,
		consumer:      consumer,
		autoAck:       autoAck,
		exclusive:     exclusive,
		preFetchCount: preFetchCount,
	}

	if err := c.init(); err != nil {
		return nil, err
	}

	return c, nil
}
