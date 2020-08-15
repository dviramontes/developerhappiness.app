-- +goose Up
-- SQL in this section is executed when the migration is applied.
create table if not exists users (
    id serial primary key,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz
);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
drop table users;
