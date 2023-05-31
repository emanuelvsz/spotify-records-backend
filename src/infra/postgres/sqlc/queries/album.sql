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