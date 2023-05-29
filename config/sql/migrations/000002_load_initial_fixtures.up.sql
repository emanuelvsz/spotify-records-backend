copy artist (id, name, super_artist_id, description, founded_at, terminated_at)
    from '/sql/sql/migrations/fixtures/000002-artist.csv'
    delimiter ';' csv header;

copy album (id, name, artist_id, release_date)
    from '/sql/sql/migrations/fixtures/000002-album.csv'
    delimiter ';' csv header;

copy song (id, name, album_id, release_date, duration)
    from '/sql/sql/migrations/fixtures/000002-song.csv'
    delimiter ';' csv header;

copy song_artist (song_id, artist_id)
    from '/sql/sql/migrations/fixtures/000002-artist_song.csv'
    delimiter ';' csv header;

copy genre (id, name, description, created_at)
    from '/sql/sql/migrations/fixtures/000002-genre.csv'
    delimiter ';' csv header;

copy artist_genre (genre_id, artist_id)
    from '/sql/sql/migrations/fixtures/000002-artist_genre.csv'
    delimiter ';' csv header;