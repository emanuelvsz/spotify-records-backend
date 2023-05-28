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

-- Tabela principal para os artistas
CREATE TABLE IF NOT EXISTS artist (
    id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(64) NOT NULL
);

-- Tabela para armazenar os gêneros
CREATE TABLE IF NOT EXISTS genre (
    id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(64) NOT NULL
);

-- Tabela de junção entre artistas e gêneros
CREATE TABLE IF NOT EXISTS artist_genre (
    artist_id UUID NOT NULL REFERENCES artist (id),
    genre_id UUID NOT NULL REFERENCES genre (id),
    PRIMARY KEY (artist_id, genre_id)
);