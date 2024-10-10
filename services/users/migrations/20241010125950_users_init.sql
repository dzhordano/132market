-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create extension if not exists "uuid-ossp";

create table if not exists users (
    id uuid  primary key default gen_random_uuid(),
    name text not null,
    email text not null,
    password text not null,
    role text not null,
    verified boolean not null,
    account_state text not null,
    account_state_since timestamp not null,
    created_at timestamp not null,
    last_seen timestamp not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table if exists users;
-- +goose StatementEnd
