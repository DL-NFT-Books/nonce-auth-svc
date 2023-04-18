-- +migrate Up

create table users (
    id bigserial primary key,
    name text not null default '',
    address text UNIQUE,
    created_at timestamp not null default CURRENT_TIMESTAMP
);

create table nonce (
    id bigserial primary key,
    message text not null,
    expiresat bigint not null,
    address Bytea

);

-- +migrate Down

drop table nonce;
drop table users;