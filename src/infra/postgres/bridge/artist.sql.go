// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: artist.sql

package bridge

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const selectArtistByID = `-- name: SelectArtistByID :one
select a.id as id,
    a.name as name,
    a.super_artist_id as super_artist_id,
    a.description as description,
    a.founded_at as founded_at,
    a.terminated_at as terminated_at,
    g.name as genre_name,
    g.description as genre_description,
    g.created_at as genre_created_at
        from artist a
    inner join artist_genre ag on a.id = ag.artist_id
    inner join genre g on g.id = ag.genre_id
    where a.id = $1
`

type SelectArtistByIDRow struct {
	ID               uuid.UUID
	Name             string
	SuperArtistID    uuid.NullUUID
	Description      sql.NullString
	FoundedAt        time.Time
	TerminatedAt     sql.NullTime
	GenreName        string
	GenreDescription sql.NullString
	GenreCreatedAt   sql.NullTime
}

func (q *Queries) SelectArtistByID(ctx context.Context, artistID uuid.UUID) (SelectArtistByIDRow, error) {
	row := q.db.QueryRowContext(ctx, selectArtistByID, artistID)
	var i SelectArtistByIDRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.SuperArtistID,
		&i.Description,
		&i.FoundedAt,
		&i.TerminatedAt,
		&i.GenreName,
		&i.GenreDescription,
		&i.GenreCreatedAt,
	)
	return i, err
}

const selectArtistSongs = `-- name: SelectArtistSongs :many
select s.id as id,
    s.name as name,
    s.album_id as album_id,
    s.release_date as release_date,
    s.duration as duration,
    sa.artist_id as artist_id
        from song s
    inner join song_artist sa on sa.song_id = s.id
    where sa.artist_id = $1
`

type SelectArtistSongsRow struct {
	ID          uuid.UUID
	Name        string
	AlbumID     uuid.NullUUID
	ReleaseDate time.Time
	Duration    string
	ArtistID    uuid.UUID
}

func (q *Queries) SelectArtistSongs(ctx context.Context, artistID uuid.UUID) ([]SelectArtistSongsRow, error) {
	rows, err := q.db.QueryContext(ctx, selectArtistSongs, artistID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SelectArtistSongsRow
	for rows.Next() {
		var i SelectArtistSongsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.AlbumID,
			&i.ReleaseDate,
			&i.Duration,
			&i.ArtistID,
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

const selectArtists = `-- name: SelectArtists :many
select a.id as id,
    a.name as name,
    a.super_artist_id as super_artist_id,
    a.description as description,
    a.founded_at as founded_at,
    a.terminated_at as terminated_at 
        from artist a
    order by a.name
`

type SelectArtistsRow struct {
	ID            uuid.UUID
	Name          string
	SuperArtistID uuid.NullUUID
	Description   sql.NullString
	FoundedAt     time.Time
	TerminatedAt  sql.NullTime
}

func (q *Queries) SelectArtists(ctx context.Context) ([]SelectArtistsRow, error) {
	rows, err := q.db.QueryContext(ctx, selectArtists)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SelectArtistsRow
	for rows.Next() {
		var i SelectArtistsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.SuperArtistID,
			&i.Description,
			&i.FoundedAt,
			&i.TerminatedAt,
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

const selectSubArtists = `-- name: SelectSubArtists :many
select a.id as id,
    a.name as name,
    a.super_artist_id as super_artist_id,
    a.description as description,
    a.founded_at as founded_at,
    a.terminated_at as terminated_at
        from artist a
    where a.super_artist_id = $1
    order by a.name
`

type SelectSubArtistsRow struct {
	ID            uuid.UUID
	Name          string
	SuperArtistID uuid.NullUUID
	Description   sql.NullString
	FoundedAt     time.Time
	TerminatedAt  sql.NullTime
}

func (q *Queries) SelectSubArtists(ctx context.Context, superArtistID uuid.NullUUID) ([]SelectSubArtistsRow, error) {
	rows, err := q.db.QueryContext(ctx, selectSubArtists, superArtistID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SelectSubArtistsRow
	for rows.Next() {
		var i SelectSubArtistsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.SuperArtistID,
			&i.Description,
			&i.FoundedAt,
			&i.TerminatedAt,
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
