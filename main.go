package main

import (
	"module/src/infra/postgres"

	"github.com/google/uuid"
)

func main() {
	postgres.NewArtistPostgresRepository().FetchArtists()
	postgres.NewSongPostgresRepository().FetchSongs()
	postgres.NewSongPostgresRepository().FetchSongsByArtistID(uuid.MustParse("7fe8fc9a-f05d-4186-8d75-447aa4dbb645"))
	postgres.NewArtistPostgresRepository().FetchArtistByID(uuid.MustParse("855f7e30-e093-4be2-bbaa-150bb23f0dd2"))
}
