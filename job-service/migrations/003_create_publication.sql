-- +goose Up
CREATE PUBLICATION job_schedules_publication FOR TABLE job_schedule;