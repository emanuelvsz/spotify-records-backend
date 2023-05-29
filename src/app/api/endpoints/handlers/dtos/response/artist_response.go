package response

import (
	"time"

	"github.com/google/uuid"
)

type ArtistDTO struct {
	ID            uuid.UUID   `json:"id"`
	Name          string      `json:"name"`
	SuperArtistID *uuid.UUID  `json:"super_artist_id,omitempty"`
	Description   *string     `json:"description"`
	FoundedAt     time.Time   `json:"founded_at"`
	TerminatedAt  *time.Time  `json:"terminated_at"`
	SubArtists    []ArtistDTO `json:"members,omitempty"`
}

func NewArtistDTO(id uuid.UUID, name string, superArtistID *uuid.UUID, description *string, foundedAt time.Time, terminatedAt *time.Time, subArtists []ArtistDTO) *ArtistDTO {
	return &ArtistDTO{
		ID:            id,
		Name:          name,
		SuperArtistID: superArtistID,
		Description:   description,
		FoundedAt:     foundedAt,
		TerminatedAt:  terminatedAt,
		SubArtists:    subArtists,
	}
}
