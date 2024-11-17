package model

import (
	"github.com/google/uuid"
	"time"
)

type Job struct {
	ID         uuid.UUID `json:"id"`
	CreateTime time.Time `json:"createTime"`
	UserID     uuid.UUID `json:"userId"`
	Recurring  bool      `json:"recurring"`
	Interval   int       `json:"interval"`
	RetryCount int       `json:"retryCount"`
}

type JobSchedule struct {
	ID         uuid.UUID `json:"id"`
	CreateTime time.Time `json:"createTime"`
	JobID      uuid.UUID `json:"jobId"`
	StartTime  time.Time `json:"startTime"`
}
