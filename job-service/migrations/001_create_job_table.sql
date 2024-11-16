-- +goose Up
CREATE TABLE IF NOT EXISTS job
(
    id          uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    create_time timestamp    not null,
    user_id     uuid         not null,
    status      varchar(255) not null,
    recurring   boolean      not null,
    interval    interval     not null,
    retry_count integer      not null
);