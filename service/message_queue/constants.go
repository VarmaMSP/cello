package message_queue

const (
	// Exchange names
	EXCHANGE_PHENOPOD_DIRECT = "phenopod_direct"
	EXCHANGE_PHENOPOD_DLX    = "phenopod_dlx"

	// Queue names
	QUEUE_IMPORT_PODCAST               = "import_podcast"
	QUEUE_REFRESH_PODCAST              = "refresh_podcast"
	QUEUE_CREATE_THUMBNAIL             = "create_thumbnail"
	QUEUE_CREATE_THUMBNAIL_DEAD_LETTER = "create_thumbnail_dead_letter"
	QUEUE_SYNC_PLAYBACK                = "sync_playback"
	QUEUE_SCHEDULED_TASK               = "scheduled_task"

	// Routing keys
	ROUTING_KEY_IMPORT_PODCAST   = "rk_import_podcast"
	ROUTING_KEY_REFRESH_PODCAST  = "rk_refresh_podcast"
	ROUTING_KEY_CREATE_THUMBNAIL = "rk_create_thumbnail"
	ROUTING_KEY_SYNC_PLAYBACK    = "rk_sync_playback"
	ROUTING_KEY_SCHEDULED_TASK   = "rk_scheduled_task"
)
