package model

const (
	JOB_NAME_CREATE_THUMBNAIL      = "create_thumbnail"
	JOB_NAME_IMPORT_PODCAST        = "import_podcast"
	JOB_NAME_REFRESH_PODCAST       = "refresh_podcast"
	JOB_NAME_SYNC_EPISODE_PLAYBACK = "sync_episode_playback"

	TASK_NAME_SCRAPE_TRENDING          = "scrape_trending"
	TASK_NAME_SCRAPE_CATEGORIES        = "scrape_categories"
	TASK_NAME_SCRAPE_ITUNES_DIRECTORY  = "scrape_itunes_directory"
	TASK_NAME_SCHEDULE_PODCAST_REFRESH = "schedule_podcast_refresh"

	QUEUE_NAME_CREATE_THUMBNAIL = "create_thumbnail"
	QUEUE_NAME_IMPORT_PODCAST   = "import_podcast"
	QUEUE_NAME_REFRESH_PODCAST  = "refresh_podcast"
	QUEUE_NAME_SYNC_PLAYBACK    = "sync_episode_playback"

	TASK_TYPE_PERIODIC  = "PERIODIC"
	TASK_TYPE_ONEOFF    = "ONEOFF"
	TASK_TYPE_IMMEDIATE = "IMMEDIATE"
)

type Job interface {
	Run()
}

type Task struct {
	Id        int64
	Name      string
	Type      string
	Interval  int
	NextRunAt int64
	Active    int
	CreatedAt int64
	UpdatedAt int64
}

func (t *Task) DbColumns() []string {
	return []string{
		"id", "name", "type_", "interval_",
		"next_run_at", "active", "created_at", "updated_at",
	}
}

func (t *Task) FieldAddrs() []interface{} {
	return []interface{}{
		&t.Id, &t.Name, &t.Type, &t.Interval,
		&t.NextRunAt, &t.Active, &t.CreatedAt, &t.UpdatedAt,
	}
}

func (t *Task) PreSave() {
	if t.CreatedAt == 0 {
		t.CreatedAt = Now()
	}

	if t.UpdatedAt == 0 {
		t.UpdatedAt = Now()
	}
}
