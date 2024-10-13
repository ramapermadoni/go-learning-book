-- +migrate Up
-- +migrate StatementBegin

create table "user" (
    id          SERIAL PRIMARY KEY,
    username    varchar(256),
    password    varchar(256),
    created_at  timestamp,
    created_by  varchar(256),
    modified_at timestamp,
    modified_by varchar(256)
)

-- +migrate StatementEnd