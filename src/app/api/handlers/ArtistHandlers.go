package handlers

import (
	"fmt"
	"module/src/core/interfaces/primary"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ArtistHandlers struct {
	service primary.ArtistManager
}

func (instance ArtistHandlers) ListArtists(context echo.Context) error {
	instance.service.ListArtists(context)
	fmt.Println("Na camada handler")
	return context.JSON(http.StatusOK, nil)
}

func NewArtistHandlers(service primary.ArtistManager) *ArtistHandlers {
	return &ArtistHandlers{service: service}
}
