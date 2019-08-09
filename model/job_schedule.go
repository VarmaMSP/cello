package model

const (
	QUEUE_NAME_SCHEDULED_JOBS = "scheduled_jobs"
	QUEUE_NAME_IMPORT_PODCAST = "import_podcast_jobs"

	JOB_NAME_CRAWL_ITUNES = "crawl_itunes"

	JOB_SCHEDULE_TYPE_PERIODIC  = "PERIODIC"
	JOB_SCHEDULE_TYPE_ONEOFF    = "ONEOFF"
	JOB_SCHEDULE_TYPE_IMMEDIATE = "IMMEDIATE"
)

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
