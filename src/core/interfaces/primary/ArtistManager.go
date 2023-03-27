package primary

import (
	"module/src/core/domain/artist"
	"module/src/core/errors"
)

type ArtistManager interface {
	ListArtists() ([]artist.Artist, errors.Error)
}
