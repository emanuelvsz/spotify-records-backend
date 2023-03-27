package postgres

import (
	"fmt"
	"log"
	"module/src/core/domain/song"
	"module/src/core/interfaces/repository"
	"module/src/infra/postgres/fixtures"

	"github.com/google/uuid"
)

type SongPostgresRepository struct{}

func (instance SongPostgresRepository) FetchSongs() ([]song.Song, error) {
	songList := fixtures.Songs
	if len(songList) == 0 {
		log.Fatal("O banco de dados está vazio, verifique se houve algum problema nele")
		return songList, nil
	}

	for _, song := range songList {
		fmt.Println(song.ID, song.Name)
	}

	return songList, nil
}

func (instance SongPostgresRepository) FetchSongsByArtistID(artistID uuid.UUID) ([]song.Song, error) {
	songList := fixtures.Songs
	var newSongList = make([]song.Song, 0)

	for _, song := range songList {
		if song.ArtistID == artistID {
			newSongList = append(newSongList, song)
		}
	}

	fmt.Println(newSongList)

	return newSongList, nil
}

func NewSongPostgresRepository() repository.SongLoader {
	return SongPostgresRepository{}
}
