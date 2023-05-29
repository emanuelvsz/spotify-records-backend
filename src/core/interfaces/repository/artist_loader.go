package repository

import (
	"module/src/core/domain/artist"
	"module/src/core/domain/song"
	"module/src/core/errors"

	"github.com/google/uuid"
)

type ArtistLoader interface {
	FindArtists() ([]artist.Artist, errors.Error)
	FindArtistSongs(artistID uuid.UUID) ([]song.Song, errors.Error)
}
