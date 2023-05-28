package response

import (
	"time"

	"github.com/google/uuid"
)

type ArtistDTO struct {
	ID            uuid.UUID   `json:"id"`
	Name          string      `json:"name"`
	SuperArtistID *uuid.UUID  `json:"superArtistId,omitempty"`
	Description   *string     `json:"description"`
	FoundedAt     time.Time   `json:"foundedAt"`
	TerminatedAt  *time.Time  `json:"terminatedAt"`
}

func NewArtistDTO(id uuid.UUID, name string, superArtistID *uuid.UUID, description *string, foundedAt time.Time, terminatedAt *time.Time) *ArtistDTO {
	return &ArtistDTO{
		ID:            id,
		Name:          name,
		SuperArtistID: superArtistID,
		Description:   description,
		FoundedAt:     foundedAt,
		TerminatedAt:  terminatedAt,
	}
}
