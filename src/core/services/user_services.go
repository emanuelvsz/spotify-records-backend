package services

import (
	"module/src/core/domain/album"
	"module/src/core/domain/artist"
	"module/src/core/domain/song"
	"module/src/core/errors"
	"module/src/core/errors/logger"
	"module/src/core/interfaces/primary"
	"module/src/core/interfaces/repository"

	"github.com/google/uuid"
)

var _ primary.UserManager = (*UserServices)(nil)

type UserServices struct {
	userRepository repository.UserLoader
	logger         logger.Logger
}

func (a UserServices) FetchAlbumSongs(albumID uuid.UUID) ([]song.Song, errors.Error) {
	songs, err := a.userRepository.FindAlbumSongs(albumID)
	if err != nil {
		a.logger.Log(err)
		return nil, err
	}

	return songs, nil
}

func (a UserServices) FetchAlbums() ([]album.Album, errors.Error) {
	albums, err := a.userRepository.FindAlbums()
	if err != nil {
		a.logger.Log(err)
		return nil, err
	}

	return albums, nil
}

func (as *UserServices) FetchArtists() ([]artist.Artist, errors.Error) {
	artists, err := as.userRepository.FindArtists()
	if err != nil {
		as.logger.Log(err)
		return nil, err
	}

	return artists, nil
}

func (as *UserServices) FetchArtistSongs(artistID uuid.UUID) ([]song.Song, errors.Error) {
	songs, err := as.userRepository.FindArtistSongs(artistID)
	if err != nil {
		as.logger.Log(err)
		return nil, err
	}

	return songs, nil
}

func (as *UserServices) FetchArtistInformation(artistID uuid.UUID) (*artist.Artist, errors.Error) {
	artist, err := as.userRepository.FindArtistInformation(artistID)
	if err != nil {
		as.logger.Log(err)
		return nil, err
	}

	return artist, nil
}

func NewUserServices(userRepository repository.UserLoader, logger logger.Logger) *UserServices {
	return &UserServices{
		userRepository: userRepository,
		logger:         logger,
	}
}
