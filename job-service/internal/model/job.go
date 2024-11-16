package model

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type Job struct {
	ID         uuid.UUID     `json:"id"`
	CreateTime time.Time     `json:"create_time"`
	UserID     uuid.UUID     `json:"user_id"`
	Status     string        `json:"status"`
	Recurring  bool          `json:"recurring"`
	Interval   time.Duration `json:"interval"`
	RetryCount int           `json:"retry_count"`
}

func (j *Job) MarshalJSON() ([]byte, error) {
	type Alias Job
	return json.Marshal(&struct {
		Alias
		Interval int `json:"interval"`
	}{
		Alias:    (Alias)(*j),
		Interval: int(j.Interval.Minutes()),
	})
}

func (j *Job) UnmarshalJSON(data []byte) error {
	type Alias Job
	aux := &struct {
		Interval int `json:"interval"`
		*Alias
	}{
		Alias: (*Alias)(j),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	j.Interval = time.Duration(aux.Interval) * time.Minute
	return nil
}
