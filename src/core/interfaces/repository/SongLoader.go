package repository

import (
	"module/src/core/domain/song"

	"github.com/google/uuid"
)

type SongLoader interface {
	FetchSongs() ([]song.Song, error)
	FetchSongsByArtistID(artistID uuid.UUID) ([]song.Song, error)
}
