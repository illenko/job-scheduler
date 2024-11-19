-- +goose Up
SELECT pg_create_logical_replication_slot('postgres_connect', 'pgoutput');