package album

import (
	"module/src/core/domain/song"

	"github.com/google/uuid"
)

type Album struct {
	ArtistID    uuid.UUID
	Name        string
	Songs       []song.Song
	ReleaseDate uuid.Time
}

