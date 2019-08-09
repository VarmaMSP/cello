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

func (s *SqlJobScheduleStore) GetAllActive() ([]*model.JobSchedule, *model.AppError) {
	m := &model.JobSchedule{}
	sql := "SELECT " + strings.Join(m.DbColumns(), ",") + " FROM job_schedule WHERE is_active = 1"

	appErrorC := model.NewAppErrorC(
		"sqlstore.sql_job_spec_store.get_all_active",
		http.StatusInternalServerError,
		nil,
	)

	rows, err := s.GetMaster().Query(sql)
	if err != nil {
		return nil, appErrorC(err.Error())
	}
	defer rows.Close()

	var res []*model.JobSchedule
	for rows.Next() {
		tmp := &model.JobSchedule{}
		if err := rows.Scan(tmp.FieldAddrs()...); err != nil {
			return nil, appErrorC(err.Error())
		}
		res = append(res, tmp)
	}
	if err := rows.Err(); err != nil {
		return nil, appErrorC(err.Error())
	}
	return res, nil
}

func (s *SqlJobScheduleStore) Disable(jobName string) *model.AppError {
	sql := `UPDATE job_schedule SET is_active = 0, updated_at = ? WHERE job_name = ?`

	_, err := s.GetMaster().Exec(sql, model.Now(), jobName)
	if err != nil {
		return model.NewAppError(
			"sqlstore.sql_job_schedule_store.disable",
			err.Error(),
			http.StatusInternalServerError,
			map[string]string{"job_name": jobName},
		)
	}
	return nil
}

func (s *SqlJobScheduleStore) SetRunAt(jobName string, runAt int64) *model.AppError {
	sql := `UPDATE job_schedule SET run_at = ?, updated_at = ? WHERE job_name = ?`

	_, err := s.GetMaster().Exec(sql, runAt, model.Now(), jobName)
	if err != nil {
		return model.NewAppError(
			"sqlstore.sql_job_schedule_store.set_run_at",
			err.Error(),
			http.StatusInternalServerError,
			map[string]string{"job_name": jobName},
		)
	}
	return nil
}
