package handlers

import (
	"module/src/core/interfaces/primary"
)

type ArtistHandlers struct {
	service primary.ArtistManager
}

func (instance ArtistHandlers) ListArtists() error {
	return nil
}
