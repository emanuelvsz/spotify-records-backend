package response

import (
	"time"

	"github.com/google/uuid"
)

type SongDTO struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	ArtistID    uuid.UUID  `json:"artist_id"`
	AlbumID     *uuid.UUID `json:"album_id,omitempty"`
	ReleaseDate time.Time  `json:"release_date"`
	Duration    string     `json:"duration"`
}

func NewSongDTO(ID uuid.UUID, Name string, ArtistID uuid.UUID, AlbumID *uuid.UUID, ReleaseDate time.Time, Duration string) *SongDTO {
	return &SongDTO{
		ID:          ID,
		Name:        Name,
		ArtistID:    ArtistID,
		AlbumID:     AlbumID,
		ReleaseDate: ReleaseDate,
		Duration:    Duration,
	}
}
