package repository

import (
	s "module/src/core/domain/song"
	"module/src/core/errors"

	"github.com/google/uuid"
)

type AlbumLoader interface {
	FindAlbumSongs(albumID uuid.UUID) ([]s.Song, errors.Error)
}
