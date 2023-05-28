package artist

import (
	"time"

	"github.com/google/uuid"
)

type Artist struct {
	id            uuid.UUID
	name          string
	superArtistID *uuid.UUID
	description   *string
	foundedAt     time.Time
	terminatedAt  *time.Time
}

func (a Artist) ID() uuid.UUID {
	return a.id
}

func (a Artist) Name() string {
	return a.name
}

func (a Artist) SuperArtistID() *uuid.UUID {
	return a.superArtistID
}

func (a Artist) Description() *string {
	return a.description
}

func (a Artist) FoundedAt() time.Time {
	return a.foundedAt
}

func (a Artist) TerminatedAt() *time.Time {
	return a.terminatedAt
}
