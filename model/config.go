package model

type Config struct {
	Mysql         MysqlConfig         `mapstructure:"mysql"`
	Rabbitmq      RabbitmqConfig      `mapstructure:"rabbitmq"`
	Elasticsearch ElasticsearchConfig `mapstrucure:"elasticsearch"`
}

// MYSQL CONFIGURATION
type MysqlConfig struct {
	Address  string `mapstructure:"address"`
	Database string `mapstructure:"database"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

// RABBITMQ CONFIGURATION
type RabbitmqConfig struct {
	Address  string               `mapstructure:"address"`
	User     string               `mapstructure:"user"`
	Password string               `mapstructure:"password"`
	Queues   RabbitmqQueuesConfig `mapstructure:"queues"`
}

type RabbitmqQueuesConfig struct {
	ScheduledJobCallQueue RabbitmqQueueConfig `mapstructure:"scheduled_job_call_queue"`
	ImportPodcastQueue    RabbitmqQueueConfig `mapstructure:"import_podcast_queue"`
	RefreshPodcastQueue   RabbitmqQueueConfig `mapstructure:"refresh_podcast_queue"`
}

type RabbitmqQueueConfig struct {
	DeliveryMode          uint8  `mapstructure:"delivery_mode"`
	ConsumerName          string `mapstructure:"consumer_name"`
	ConsumerAutoAck       bool   `mapstructure:"consumer_auto_ack"`
	ConsumerExclusive     bool   `mapstructure:"consumer_exclusive"`
	ConsumerPreFetchCount int    `mapstructure:"consumer_prefetch_count"`
	ConsumerWorkerLimit   int    `mspstructure:"consumer_worker_limit"`
}

// ELASTICSEARCH CONFIGURATION
type ElasticsearchConfig struct {
	Address  string `mapstructure:"address"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}
