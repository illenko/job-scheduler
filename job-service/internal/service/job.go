package service

import (
	"context"
	"encoding/json"
	"job-service/internal/model"
	"job-service/internal/repository"
	"log"
	"time"
)

type JobService struct {
	repo            *repository.JobRepository
	jobScheduleRepo *repository.JobScheduleRepository
}

func NewJobService(repo *repository.JobRepository, jobScheduleRepo *repository.JobScheduleRepository) *JobService {
	return &JobService{
		repo:            repo,
		jobScheduleRepo: jobScheduleRepo,
	}
}

func (s *JobService) GetJobs(ctx context.Context) ([]model.Job, error) {
	log.Println("Fetching jobs from the repository")
	jobs, err := s.repo.GetJobs(ctx)
	if err != nil {
		log.Printf("Error fetching jobs: %v", err)
		return nil, err
	}
	log.Printf("Successfully fetched %d jobs", len(jobs))
	return jobs, nil
}

func (s *JobService) CreateJob(ctx context.Context, job model.Job) (model.Job, error) {
	jobData, _ := json.Marshal(job)
	log.Printf("Creating a new job with details: %s", jobData)
	err := s.repo.CreateJob(ctx, &job)
	if err != nil {
		log.Printf("Error creating job: %v", err)
		return model.Job{}, err
	}
	log.Println("Successfully created job")

	jobSchedule := model.JobSchedule{
		JobID:     job.ID,
		StartTime: job.CreateTime.Add(time.Duration(job.Interval) * time.Second),
	}

	err = s.jobScheduleRepo.CreateJobSchedule(ctx, &jobSchedule)
	if err != nil {
		log.Printf("Error creating job schedule: %v", err)
		return model.Job{}, err
	}
	log.Println("Successfully created job schedule")
	return job, nil
}
