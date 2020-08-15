-- +goose Up
-- SQL in this section is executed when the migration is applied.
alter table users
    add column name varchar(32),
    add column active boolean default false,
    add column is_bot boolean default false,
    add column email varchar(255),
    add column timezone varchar(128),
    add column img_url varchar(255),
    add column is_admin boolean default false,
    add column is_owner boolean default false;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
alter table users
    drop column name,
    drop column active,
    drop column is_bot,
    drop column email,
    drop column timezone,
    drop column img_url,
    drop column is_admin,
    drop column is_owner;
