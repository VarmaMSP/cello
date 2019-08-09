package model

const (
	JOB_NAME_SCRAPE_ITUNES  = "scrape_itunes"
	JOB_NAME_IMPORT_PODCAST = "import_podcast"

	QUEUE_NAME_SCHEDULED_WORK = "scheduled_work"
	QUEUE_NAME_IMPORT_PODCAST = "import_podcast"
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
