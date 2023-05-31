package services

import (
	s "module/src/core/domain/song"
	"module/src/core/errors"
	"module/src/core/errors/logger"
	"module/src/core/interfaces/primary"
	"module/src/core/interfaces/repository"

	"github.com/google/uuid"
)

var _ primary.AlbumManager = (*AlbumServices)(nil)

type AlbumServices struct {
	albumRepository repository.AlbumLoader
	logger          logger.Logger
}

func (a AlbumServices) FetchAlbumSongs(albumID uuid.UUID) ([]s.Song, errors.Error) {
	songs, err := a.albumRepository.FindAlbumSongs(albumID)
	if err != nil {
		a.logger.Log(err)
		return nil, err
	}

	return songs, nil
}

func NewAlbumServices(albumRepository repository.AlbumLoader, logger logger.Logger) *AlbumServices {
	return &AlbumServices{
		albumRepository: albumRepository,
		logger:          logger,
	}
}
