-- +goose Up
create table if not exists title
(
    code    varchar primary key,
    name_ru varchar
);

-- +goose Down
drop table if exists title;
