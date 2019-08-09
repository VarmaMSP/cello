package model

const (
	JOB_SCRAPE_ITUNES  = "scrape_itunes"
	JOB_IMPORT_PODCAST = "import_podcast"
)

// A job takes input and does some work with it
type Job interface {
	Start() *AppError
	Stop() *AppError
	InputChan() chan interface{}
}

type ImportPodcastInput struct {
	Id      string `json:"id"`
	Source  string `json:"source"`
	FeedUrl string `json:"feed_url"`
}
