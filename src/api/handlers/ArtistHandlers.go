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

func NewArtistHandlers(service primary.ArtistManager) *ArtistHandlers {
	return &ArtistHandlers{service: service}
}
