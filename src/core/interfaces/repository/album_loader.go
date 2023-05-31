package repository

import (
	"module/src/core/domain/album"
	"module/src/core/domain/song"
	"module/src/core/errors"

	"github.com/google/uuid"
)

type AlbumLoader interface {
	FindAlbums() ([]album.Album, errors.Error)
	FindAlbumSongs(albumID uuid.UUID) ([]song.Song, errors.Error)
}
