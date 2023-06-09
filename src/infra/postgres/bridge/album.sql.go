// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: album.sql

package bridge

import (
	"context"

	"github.com/google/uuid"
)

const selectAlbumSongs = `-- name: SelectAlbumSongs :many
select s.id as id,
    s.name as name,
    s.album_id as album_id,
    s.release_date as release_date,
    s.duration as duration,
    s.lyrics as lyrics,
    s.track_number as track_number,
    s.spotify_url as spotify_url
    from song s where s.album_id = $1
    order by s.track_number
`

func (q *Queries) SelectAlbumSongs(ctx context.Context, albumID uuid.NullUUID) ([]Song, error) {
	rows, err := q.db.QueryContext(ctx, selectAlbumSongs, albumID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Song
	for rows.Next() {
		var i Song
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.AlbumID,
			&i.ReleaseDate,
			&i.Duration,
			&i.Lyrics,
			&i.TrackNumber,
			&i.SpotifyUrl,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectAlbums = `-- name: SelectAlbums :many
select a.id as id,
    a.name as name,
    a.artist_id as artist_id,
    a.release_date as release_date,
    a.description as description,
    a.image_url as image_url
    from album a
`

func (q *Queries) SelectAlbums(ctx context.Context) ([]Album, error) {
	rows, err := q.db.QueryContext(ctx, selectAlbums)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Album
	for rows.Next() {
		var i Album
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.ArtistID,
			&i.ReleaseDate,
			&i.Description,
			&i.ImageUrl,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
