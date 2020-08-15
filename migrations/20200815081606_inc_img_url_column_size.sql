-- +goose Up
-- SQL in this section is executed when the migration is applied.
alter table users alter column img_url type varchar(510);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
alter table users alter column img_url type varchar(255);
