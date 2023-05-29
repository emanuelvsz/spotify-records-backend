package services

import (
	a "module/src/core/domain/artist"
	"module/src/core/domain/song"
	"module/src/core/errors"
	"module/src/core/errors/logger"
	"module/src/core/interfaces/primary"
	"module/src/core/interfaces/repository"

	"github.com/google/uuid"
)

var _ primary.ArtistManager = (*ArtistServices)(nil)

type ArtistServices struct {
	artistRepository repository.ArtistLoader
	logger           logger.Logger
}

func (as *ArtistServices) FetchArtists() ([]a.Artist, errors.Error) {
	artists, err := as.artistRepository.FindArtists()
	if err != nil {
		as.logger.Log(err)
		return nil, err
	}

	return artists, nil
}

func (as *ArtistServices) FetchArtistSongs(artistID uuid.UUID) ([]song.Song, errors.Error) {
	songs, err := as.artistRepository.FindArtistSongs(artistID)
	if err != nil {
		as.logger.Log(err)
		return nil, err
	}

	return songs, nil
}

func (as *ArtistServices) FetchArtistInformation(artistID uuid.UUID) (*a.Artist, errors.Error) {
	songs, err := as.artistRepository.FindArtistInformation(artistID)
	if err != nil {
		as.logger.Log(err)
		return nil, err
	}

	return songs, nil
}

func NewArtistServices(artistRepository repository.ArtistLoader, logger logger.Logger) *ArtistServices {
	return &ArtistServices{
		artistRepository: artistRepository,
		logger:           logger,
	}
}
