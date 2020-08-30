package model

import "github.com/varmamsp/cello/util/datetime"

const (
	TASK_NAME_SCRAPE_TRENDING          = "scrape_trending"
	TASK_NAME_SCRAPE_CATEGORIES        = "scrape_categories"
	TASK_NAME_SCRAPE_ITUNES_DIRECTORY  = "scrape_itunes_directory"
	TASK_NAME_SCHEDULE_PODCAST_REFRESH = "schedule_podcast_refresh"
	TASK_NAME_REIMPORT_PODCASTS        = "reimport_podcasts"
	TASK_NAME_INDEX_EPISODES           = "reindex_episodes"
	TASK_NAME_INDEX_PODCASTS           = "reindex_podcasts"
	TASK_NAME_INDEX_KEYWORDS           = "index_keywords"
	TASK_NAME_FIX_CATEGORIES           = "fix_categories"
	TASK_NAME_EXTRACT_KEYWORDS         = "extract_keywords"
	TASK_NAME_FIX_KEYWORDS             = "fix_keywords"
	TASK_NAME_CLEAN_UP_KEYWORDS        = "clean_up_keywords"

	TASK_TYPE_PERIODIC  = "PERIODIC"
	TASK_TYPE_ONEOFF    = "ONEOFF"
	TASK_TYPE_IMMEDIATE = "IMMEDIATE"
)

type Job interface {
	Run()
}

type Task struct {
	Id        int64  `db:"id"`
	Name      string `db:"name"`
	Type      string `db:"type"`
	Interval  int    `db:"interval"`
	NextRunAt int64  `db:"next_run_at"`
	Active    int    `db:"active"`
	CreatedAt int64  `db:"created_at"`
	UpdatedAt int64  `db:"updated_at"`
}

func (t *Task) DbColumns() []string {
	return []string{
		"id", "name", "type", "interval_",
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
		t.CreatedAt = datetime.Unix()
	}

	if t.UpdatedAt == 0 {
		t.UpdatedAt = datetime.Unix()
	}
}
