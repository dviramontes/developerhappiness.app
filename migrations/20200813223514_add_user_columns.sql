-- +goose Up
-- SQL in this section is executed when the migration is applied.
alter table users
    add column name varchar(32),
    add column active boolean default false,
    add column isBot boolean default false,
    add column email varchar(255),
    add column timezone varchar(128),
    add column imgUrl varchar(255),
    add column isAdmin boolean default false,
    add column isOwner boolean default false;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
alter table users
    drop column name,
    drop column active,
    drop column isBot,
    drop column email,
    drop column timezone,
    drop column imgUrl,
    drop column isAdmin,
    drop column isOwner;
