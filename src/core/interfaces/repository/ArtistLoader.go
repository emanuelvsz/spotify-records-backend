package repository

import (
	"module/src/core/domain/artist"
	"module/src/core/errors"
)

type ArtistLoader interface {
	FetchArtists() ([]artist.Artist, errors.Error)
}
