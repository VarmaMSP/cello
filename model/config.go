package model

type Config struct {
	Mysql         Mysql         `mapstructure:"mysql"`
	Rabbitmq      Rabbitmq      `mapstructure:"rabbitmq"`
	Elasticsearch Elasticsearch `mapstructure:"elasticsearch"`
	Queues        Queues        `mapstructure:"queues"`
	Jobs          Jobs          `mapstructure:"jobs"`
	OAuth         OAuth         `mapstructure:"oauth"`
}

// MYSQL CONFIGURATION
type Mysql struct {
	Address  string `mapstructure:"address"`
	Database string `mapstructure:"database"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

// RABBITMQ CONFIGURATION
type Rabbitmq struct {
	Address  string `mapstructure:"address"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

// ELASTICSEARCH CONFIGURATION
type Elasticsearch struct {
	Address  string `mapstructure:"address"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

// Queue
type Queues struct {
	ImportPodcast   Queue_ `mapstructure:"import_podcast"`
	RefreshPodcast  Queue_ `mapstructure:"refresh_podcast"`
	CreateThumbnail Queue_ `mapstructure:"create_thumbnail"`
}

type Queue_ struct {
	DeliveryMode          uint8  `mapstructure:"delivery_mode"`
	ConsumerName          string `mapstructure:"consumer_name"`
	ConsumerAutoAck       bool   `mapstructure:"consumer_auto_ack"`
	ConsumerExclusive     bool   `mapstructure:"consumer_exclusive"`
	ConsumerPreFetchCount int    `mapstructure:"consumer_prefetch_count"`
}

// Job
type Jobs struct {
	Scheduler       Job_ `mapstructure:"scheduler"`
	ImportPodcast   Job_ `mapstructure:"import_podcast"`
	RefreshPodcast  Job_ `mapstructure:"refresh_podcast"`
	CreateThumbnail Job_ `mapstructure:"create_thumbnail"`
}

type Job_ struct {
	Enable      bool `mapstructure:"enable"`
	WorkerLimit int  `mapstructure:"worker_limit"`
}

// OAuth
type OAuth struct {
	Google   OAuth_ `mapstructure:"google"`
	Facebook OAuth_ `mapstructure:"facebook"`
}

type OAuth_ struct {
	ClientId     string   `mapstructure:"client_id"`
	ClientSecret string   `mapstructure:"client_secret"`
	RedirectUrl  string   `mapstructure:"redirect_url"`
	Scopes       []string `mapstructure:"scopes"`
}
