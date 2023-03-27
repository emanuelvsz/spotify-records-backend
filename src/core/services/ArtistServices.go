package services

import (
	"fmt"
	"log"
	"module/src/core/domain/artist"
	"module/src/core/interfaces/repository"

	"github.com/google/uuid"
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

func (instance ArtistServices) FetchArtistByID(artistID uuid.UUID) ([]artist.Artist, error) {
	artistList, err := instance.artistRepository.FetchArtists()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Println("Na camada handler")
	fmt.Println(artistList)

	return artistList, nil
}

func NewArtistServices() ArtistServices {
	return ArtistServices{}
}
