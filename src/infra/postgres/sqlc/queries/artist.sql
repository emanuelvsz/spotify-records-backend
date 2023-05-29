-- name: SelectArtists :many
select a.id as id,
    a.name as name,
    a.super_artist_id as super_artist_id,
    a.description as description,
    a.founded_at as founded_at,
    terminated_at as terminated_at 
        from artist a
    order by a.name;

-- name: SelectArtistSongs :many
select s.id as id,
    s.name as name,
    s.album_id as album_id,
    s.release_date as release_date,
    s.duration as duration,
    sa.artist_id as artist_id
        from song s
    inner join song_artist sa on sa.song_id = s.id
    where sa.artist_id = @artist_id;