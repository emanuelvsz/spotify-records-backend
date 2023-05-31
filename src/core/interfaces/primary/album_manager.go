package primary

import (
	s "module/src/core/domain/song"
	"module/src/core/errors"

	"github.com/google/uuid"
)

type AlbumManager interface {
	FetchAlbumSongs(albumID uuid.UUID) ([]s.Song, errors.Error)
}
