-- +migrate Up
-- +migrate StatementBegin

create table kategori (
    id          SERIAL PRIMARY KEY,
    name        varchar(256),
    created_at  timestamp,
    created_by  varchar(256),
    modified_at timestamp,
    modified_by varchar(256)
)

-- +migrate StatementEnd