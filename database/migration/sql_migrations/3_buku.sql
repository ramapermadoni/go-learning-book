-- +migrate Up
-- +migrate StatementBegin

create table buku (
    id              SERIAL PRIMARY KEY,
    title           varchar(256),
    description     varchar,
    image_url       varchar(256),
    release_year    int,
    price           int,
    total_page      int,
    thickness       varchar(15),
    category_id     int REFERENCES kategori(id), -- Mendefinisikan foreign key
    created_at      timestamp,
    created_by      varchar(256),
    modified_at     timestamp,
    modified_by     varchar(256)
)

-- +migrate StatementEnd