package primary

import (
	"module/src/core/domain/artist"
	"module/src/core/errors"
)

type ArtistManager interface {
	FetchArtists() ([]artist.Artist, errors.Error)
}
