package primary

import (
	"module/src/core/domain/album"
	"module/src/core/domain/song"
	"module/src/core/errors"

	"github.com/google/uuid"
)

type AlbumManager interface {
	FetchAlbums() ([]album.Album, errors.Error)
	FetchAlbumSongs(albumID uuid.UUID) ([]song.Song, errors.Error)
}
