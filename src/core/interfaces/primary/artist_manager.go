package primary

import (
	"module/src/core/domain/artist"
	"module/src/core/domain/song"
	"module/src/core/errors"

	"github.com/google/uuid"
)

type ArtistManager interface {
	FetchArtists() ([]artist.Artist, errors.Error)
	FetchArtistSongs(artistID uuid.UUID) ([]song.Song, errors.Error)
}
