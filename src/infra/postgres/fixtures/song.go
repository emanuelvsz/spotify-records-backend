package fixtures

import (
	"module/src/core/domain/song"

	"github.com/google/uuid"
)

var Songs = []song.Song{
	{
		ID:       uuid.New(),
		ArtistID: uuid.MustParse("855f7e30-e093-4be2-bbaa-150bb23f0dd2"),
		Name:     "As it was",
	},
	{
		ID:       uuid.New(),
		ArtistID: uuid.MustParse("e5584dbc-2aa5-4a1d-b03b-186a9bd21405"),
		Name:     "Lover",
	},
	{
		ID:       uuid.New(),
		ArtistID: uuid.MustParse("7fe8fc9a-f05d-4186-8d75-447aa4dbb645"),
		Name:     "Run",
	},
	{
		ID:       uuid.New(),
		ArtistID: uuid.MustParse("7fe8fc9a-f05d-4186-8d75-447aa4dbb645"),
		Name:     "Sanctuary",
	},
	{
		ID:       uuid.New(),
		ArtistID: uuid.MustParse("af14b280-a023-4d15-a70c-3345e71228e7"),
		Name:     "Break my heart again",
	},
	{
		ID:       uuid.New(),
		ArtistID: uuid.MustParse("af14b280-a023-4d15-a70c-3345e71228e7"),
		Name:     "Lets fall in love for the night",
	},
}
