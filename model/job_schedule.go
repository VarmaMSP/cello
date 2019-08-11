package model

const (
	JOB_NAME_SCRAPE_ITUNES   = "scrape_itunes"
	JOB_NAME_IMPORT_PODCAST  = "import_podcast"
	JOB_NAME_REFRESH_PODCAST = "refresh_podcast"

	QUEUE_NAME_SCHEDULED_WORK  = "scheduled_work"
	QUEUE_NAME_IMPORT_PODCAST  = "import_podcast"
	QUEUE_NAME_REFRESH_PODCAST = "refresh_podcast"

	JOB_SCHEDULE_TYPE_PERIODIC  = "PERIODIC"
	JOB_SCHEDULE_TYPE_ONEOFF    = "ONEOFF"
	JOB_SCHEDULE_TYPE_IMMEDIATE = "IMMEDIATE"
)

// A job takes input and does some work with it
type Job interface {
	Start() *AppError
	Stop() *AppError
	InputChan() chan interface{}
}

type ScheduledWorkInput struct {
	JobName string `json:"job_name"`
}

type ImportPodcastInput struct {
	Id      string `json:"id"`
	Source  string `json:"source"`
	FeedUrl string `json:"feed_url"`
}

type RefreshPodcastInput struct {
	PodcastId        string `json:"podcast_id"`
	FeedUrl          string `json:"feed_url"`
	FeedEtag         string `json:"feed_etag"`
	FeedLastModified string `json:"feed_last_modified"`
}

type JobSchedule struct {
	JobName   string
	Type      string
	RunAt     int64
	RunAfter  int64
	IsActive  int
	CreatedAt int64
	UpdatedAt int64
}

func (j *JobSchedule) DbColumns() []string {
	return []string{
		"job_name", "type", "run_at", "run_after",
		"is_active", "created_at", "updated_at",
	}
}

func (j *JobSchedule) FieldAddrs() []interface{} {
	var i []interface{}
	return append(i,
		&j.JobName, &j.Type, &j.RunAt, &j.RunAfter,
		&j.IsActive, &j.CreatedAt, &j.UpdatedAt,
	)
}
