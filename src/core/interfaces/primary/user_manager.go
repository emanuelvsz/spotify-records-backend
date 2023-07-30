package primary

import (
	"module/src/core/domain/album"
	"module/src/core/domain/artist"
	"module/src/core/domain/song"
	"module/src/core/errors"

	"github.com/google/uuid"
)

type UserManager interface {
	FetchAlbums() ([]album.Album, errors.Error)
	FetchAlbumSongs(albumID uuid.UUID) ([]song.Song, errors.Error)
	FetchArtists() ([]artist.Artist, errors.Error)
	FetchArtistSongs(artistID uuid.UUID) ([]song.Song, errors.Error)
	FetchArtistInformation(artistID uuid.UUID) (*artist.Artist, errors.Error)
}
