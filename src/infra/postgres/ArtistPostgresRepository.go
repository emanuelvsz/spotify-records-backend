package postgres

import (
	"fmt"
	"log"
	"module/src/core/domain/artist"
	"module/src/core/interfaces/repository"
	"module/src/infra/postgres/fixtures"

	"github.com/google/uuid"
)

type ArtistPostgresRepository struct{}

func (instance ArtistPostgresRepository) FetchArtists() ([]artist.Artist, error) {
	artistList := fixtures.Artists
	if len(artistList) == 0 {
		log.Fatal("O banco de dados está vazio, verifique se houve algum problema nele")
		return artistList, nil
	}

	for _, artist := range artistList {
		fmt.Println(artist.ID, artist.Name)
	}

	return artistList, nil
}

func (instance ArtistPostgresRepository) FetchArtistByID(artistID uuid.UUID) (artist.Artist, error) {
	artistList := fixtures.Artists
	var newArtist artist.Artist

	for _, artist := range artistList {
		if artist.ID == artistID {
			newArtist = artist
		}
	}

	fmt.Println(newArtist)

	return newArtist, nil
}

func NewArtistPostgresRepository() repository.ArtistLoader {
	return ArtistPostgresRepository{}
}
