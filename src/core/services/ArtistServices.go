package services

import (
	"log"
	"module/src/core/domain/artist"
	"module/src/core/interfaces/repository"
)

type ArtistServices struct {
	artistRepository repository.ArtistLoader
}

func (instance ArtistServices) FetchArtists() ([]artist.Artist, error) {
	artistList, err := instance.artistRepository.FetchArtists()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return artistList, nil
}

func NewArtistServices() ArtistServices {
	return ArtistServices{}
}