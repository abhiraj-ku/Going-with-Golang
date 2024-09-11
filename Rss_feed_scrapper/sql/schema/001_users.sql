-- +goose Up

create table users(
    id UUID,
    created_at timestamp not null,
    updated_at timestamp not null,
    name text not null
);

-- +goose Down
drop table users;
