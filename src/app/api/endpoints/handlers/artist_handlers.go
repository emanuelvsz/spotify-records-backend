package handlers

import (
	"module/src/core/interfaces/primary"
)

type ArtistHandlers struct {
	service primary.ArtistManager
}

func NewArtistHandlers(service primary.ArtistManager) *ArtistHandlers {
	return &ArtistHandlers{service: service}
}
