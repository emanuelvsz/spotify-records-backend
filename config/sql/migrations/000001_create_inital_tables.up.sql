create extension if not exists "uuid-ossp";

create table if not exists role
(
    id      uuid            not null
        constraint pk_role_id primary key,
    name    varchar         not null,
    code    varchar         not null
);

create table if not exists account
(
    id          uuid        not null
        constraint pk_account_id primary key
        constraint df_account_id default uuid_generate_v4(),
    role_id     uuid        not null,
    name        varchar(64) not null,
    email       varchar(64) not null,
    hash        varchar(64) not null,
    password    text        not null


    constraint fk_account_role_id
        foreign key (role_id) references role (id)
);
