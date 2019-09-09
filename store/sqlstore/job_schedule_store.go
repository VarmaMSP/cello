package sqlstore

import (
	"net/http"
	"strings"

	"github.com/varmamsp/cello/model"
)

type SqlJobScheduleStore struct {
	SqlStore
}

func NewSqlJobScheduleStore(store SqlStore) *SqlJobScheduleStore {
	return &SqlJobScheduleStore{store}
}

func (s *SqlJobScheduleStore) GetAllActive() (res []*model.JobSchedule, appE *model.AppError) {
	sql := "SELECT " + strings.Join((&model.JobSchedule{}).DbColumns(), ",") + " FROM job_schedule WHERE is_active = 1"

	copyTo := func() []interface{} {
		tmp := &model.JobSchedule{}
		res = append(res, tmp)
		return tmp.FieldAddrs()
	}

	if err := s.Query(copyTo, sql); err != nil {
		appE = model.NewAppError(
			"sqlstore.sql_job_spec_store.get_all_active", err.Error(), http.StatusInternalServerError,
			nil,
		)
	}
	return
}

func (s *SqlJobScheduleStore) Disable(jobName string) *model.AppError {
	sql := `UPDATE job_schedule SET is_active = 0, updated_at = ? WHERE job_name = ?`

	if _, err := s.GetMaster().Exec(sql, model.Now(), jobName); err != nil {
		return model.NewAppError(
			"sqlstore.sql_job_schedule_store.disable", err.Error(), http.StatusInternalServerError,
			map[string]string{"job_name": jobName},
		)
	}
	return nil
}

func (s *SqlJobScheduleStore) SetRunAt(jobName string, runAt int64) *model.AppError {
	sql := `UPDATE job_schedule SET run_at = ?, updated_at = ? WHERE job_name = ?`

	if _, err := s.GetMaster().Exec(sql, runAt, model.Now(), jobName); err != nil {
		return model.NewAppError(
			"sqlstore.sql_job_schedule_store.set_run_at", err.Error(), http.StatusInternalServerError,
			map[string]string{"job_name": jobName},
		)
	}
	return nil
}
