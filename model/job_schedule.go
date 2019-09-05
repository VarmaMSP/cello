package model

import (
	"github.com/streadway/amqp"
)

const (
	JOB_NAME_SCRAPE_ITUNES    = "scrape_itunes"
	JOB_NAME_IMPORT_PODCAST   = "import_podcast"
	JOB_NAME_REFRESH_PODCAST  = "refresh_podcast"
	JOB_NAME_SCHEDULE_REFRESH = "schedule_refresh"

	QUEUE_NAME_SCHEDULED_JOB_CALL = "scheduled_work"
	QUEUE_NAME_IMPORT_PODCAST     = "import_podcast"
	QUEUE_NAME_REFRESH_PODCAST    = "refresh_podcast"
	QUEUE_NAME_CREATE_THUMBNAIL   = "create_thumbnail"

	JOB_SCHEDULE_TYPE_PERIODIC  = "PERIODIC"
	JOB_SCHEDULE_TYPE_ONEOFF    = "ONEOFF"
	JOB_SCHEDULE_TYPE_IMMEDIATE = "IMMEDIATE"
)

// A job takes input and does some work with it
type Job interface {
	Call(delivery amqp.Delivery)
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
