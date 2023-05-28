create extension if not exists "uuid-ossp";

create table if not exists user_account
(
    id    uuid          not null
        constraint pk_user_account_id primary key,
    pwd   text          not null,
    hash  varchar(1024) not null,
    email varchar(64)   not null
);