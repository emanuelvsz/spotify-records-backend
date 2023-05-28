copy user_account (id, pwd, hash, email)
    from '/sql/sql/migrations/fixtures/000002-user_account.csv'
    delimiter ';' csv header;

copy artist (id, name)
    from '/sql/sql/migrations/fixtures/000001-artist.csv'
    delimiter ';' csv header;

copy genre (id, name)
    from '/sql/sql/migrations/fixtures/000001-genre.csv'
    delimiter ';' csv header;

copy artist_genre (genre_id, artist_id)
    from '/sql/sql/migrations/fixtures/000001-artist_genre.csv'
        delimiter ';' csv header;
