-- +goose Up
CREATE TABLE IF NOT EXISTS job
(
    id          uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    create_time timestamp not null,
    user_id     uuid      not null,
    recurring   boolean   not null,
    interval    integer   not null,
    retry_count integer   not null
);
