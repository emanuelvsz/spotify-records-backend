-- name: SelectArtists :many
select a.id as id,
    a.name as name,
    a.super_artist_id as super_artist_id,
    a.description as description,
    a.founded_at as founded_at,
    terminated_at as terminated_at 
        from artist a
    order by a.name;