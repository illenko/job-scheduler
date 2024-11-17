-- +goose Up
CREATE TABLE IF NOT EXISTS job_schedule
(
    id          uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    create_time timestamp not null,
    job_id      uuid      not null,
    start_time  timestamp not null
);
