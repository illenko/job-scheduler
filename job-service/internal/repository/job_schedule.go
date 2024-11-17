package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"job-service/internal/model"
)

type JobScheduleRepository struct {
	dbpool *pgxpool.Pool
}

func NewJobScheduleRepository(dbpool *pgxpool.Pool) *JobScheduleRepository {
	return &JobScheduleRepository{dbpool: dbpool}
}

func (r *JobScheduleRepository) CreateJobSchedule(ctx context.Context, jobSchedule *model.JobSchedule) error {
	err := r.dbpool.QueryRow(ctx,
		"INSERT INTO job_schedule (create_time, job_id, start_time) VALUES (NOW(), $1, $2) RETURNING id, create_time",
		jobSchedule.JobID, jobSchedule.StartTime).Scan(&jobSchedule.ID, &jobSchedule.CreateTime)
	if err != nil {
		return err
	}
	return nil
}
