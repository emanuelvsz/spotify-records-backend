package fixtures

import (
	"module/src/core/domain/artist"

	"github.com/google/uuid"
)

var Artists = []artist.Artist{
	{
		ID:          uuid.MustParse("855f7e30-e093-4be2-bbaa-150bb23f0dd2"),
		Name:        "Harry Styles",
		Age:         28,
		Nacionality: "American",
	},
	{
		ID:          uuid.MustParse("7fe8fc9a-f05d-4186-8d75-447aa4dbb645"),
		Name:        "Joji",
		Age:         30,
		Nacionality: "Japanese",
	},
	{
		ID:          uuid.MustParse("e5584dbc-2aa5-4a1d-b03b-186a9bd21405"),
		Name:        "Taylor Swift",
		Age:         28,
		Nacionality: "American",
	},
	{
		ID:          uuid.MustParse("af14b280-a023-4d15-a70c-3345e71228e7"),
		Name:        "FINNEAS",
		Age:         28,
		Nacionality: "American",
	},
}
