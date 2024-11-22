package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"job-service/internal/model"
)

type JobRepository struct {
	dbpool *pgxpool.Pool
}

func NewJobRepository(dbpool *pgxpool.Pool) *JobRepository {
	return &JobRepository{dbpool: dbpool}
}

func (r *JobRepository) GetJobs(ctx context.Context) ([]model.Job, error) {
	rows, err := r.dbpool.Query(ctx, "SELECT id, create_time, user_id, recurring, interval, retry_count FROM job")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobs []model.Job
	for rows.Next() {
		var job model.Job
		if err := rows.Scan(&job.ID, &job.CreateTime, &job.UserID, &job.Recurring, &job.Interval, &job.RetryCount); err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}
	return jobs, nil
}

func (r *JobRepository) CreateJob(ctx context.Context, job *model.Job) error {
	err := r.dbpool.QueryRow(ctx,
		"INSERT INTO job (create_time, user_id, recurring, interval, retry_count) VALUES (NOW(), $1, $2, $3, $4) RETURNING id, create_time",
		job.UserID, job.Recurring, job.Interval, job.RetryCount).Scan(&job.ID, &job.CreateTime)
	if err != nil {
		return err
	}
	return nil
}
