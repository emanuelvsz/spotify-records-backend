package repository

import (
	"module/src/core/domain/album"
	"module/src/core/domain/artist"
	"module/src/core/domain/song"
	"module/src/core/errors"

	"github.com/google/uuid"
)

type UserLoader interface {
	FindAlbums() ([]album.Album, errors.Error)
	FindAlbumSongs(albumID uuid.UUID) ([]song.Song, errors.Error)
	FindArtists() ([]artist.Artist, errors.Error)
	FindArtistSongs(artistID uuid.UUID) ([]song.Song, errors.Error)
	FindArtistInformation(artistID uuid.UUID) (*artist.Artist, errors.Error)
}
