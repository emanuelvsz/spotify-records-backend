package repository

import (
	"module/src/core/domain/artist"

	"github.com/google/uuid"
)

type ArtistLoader interface {
	FetchArtists() ([]artist.Artist, error)
	FetchArtistByID(artistID uuid.UUID) (artist.Artist, error)
}
