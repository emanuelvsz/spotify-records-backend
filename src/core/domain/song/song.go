package song

import (
	"time"

	"github.com/google/uuid"
)

type Song struct {
	id          uuid.UUID
	name        string
	artistID    uuid.UUID
	albumID     *uuid.UUID
	releaseDate time.Time
	duration    float32
}

func (s Song) ID() uuid.UUID {
	return s.id
}

func (s Song) Name() string {
	return s.name
}

func (s Song) ArtistID() uuid.UUID {
	return s.artistID
}

func (s Song) AlbumID() *uuid.UUID {
	return s.albumID
}

func (s Song) ReleaseDate() time.Time {
	return s.releaseDate
}

func (s Song) Duration() float32 {
	return s.duration
}
