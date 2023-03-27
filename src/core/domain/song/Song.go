package song

import "github.com/google/uuid"

type Song struct {
	ID       uuid.UUID
	ArtistID uuid.UUID
	Name     string
}
