copy artist (id, name, super_artist_id, description, founded_at, terminated_at)
    from '/sql/sql/migrations/fixtures/000002-artist.csv'
    delimiter ';' csv header;
