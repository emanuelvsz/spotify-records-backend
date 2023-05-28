create extension if not exists "uuid-ossp";

create table if not exists user_account
(
    id    uuid          not null
        constraint pk_user_account_id primary key,
    pwd   text          not null,
    hash  varchar(1024) not null,
    email varchar(64)   not null
);

create table if not exists profile
(
    id          uuid        not null
        constraint pk_profile_id primary key,
    user_id     uuid        not null,
    name        varchar(64) not null,
    picture_url varchar(64) null,


    constraint fk_profile_user_account
        foreign key (user_id) references user_account (id)
);