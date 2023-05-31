-- name: SelectAlbumSongs :many
select s.id as id,
    s.name as name,
    s.album_id as album_id,
    s.release_date as release_date,
    s.duration as duration,
    s.lyrics as lyrics,
    s.track_number as track_number,
    s.spotify_url as spotify_url
    from song s where s.album_id = @album_id
    order by s.track_number;

-- name: SelectAlbums :many
select a.id as id,
    a.name as name,
    a.artist_id as artist_id,
    a.release_date as release_date,
    a.description as description,
    a.image_url as image_url
    from album a;