package handlers

import (
	"module/src/app/api/endpoints/handlers/dtos/response"
	"module/src/core/interfaces/primary"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ArtistHandlers struct {
	service primary.ArtistManager
}

// GetArtists
// @ID GetArtists
// @Summary Buscar todos os artistas do banco
// @Tags Rotas do usuário
// @Description Rota que permite que se busque todos os artistas do banco
// @Produce json
// @Success 200 {array} response.ArtistDTO "Requisição realizada com sucesso."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 422 {object} response.ErrorMessage "Algum dado informado não pôde ser processado."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado."
// @Failure 503 {object} response.ErrorMessage "A base de dados não está disponível."
// @Router /artists [get]
func (h ArtistHandlers) GetArtists(context echo.Context) error {
	artistsSlice, fetchErr := h.service.FetchArtists()
	if fetchErr != nil {
		return getHttpHandledErrorResponse(context, fetchErr)
	}

	var artists []response.ArtistDTO
	for _, each := range artistsSlice {
		superArtistID := each.SuperArtistID()
		if *each.SuperArtistID() == uuid.Nil {
			superArtistID = nil
		}

		newArtist := *response.NewArtistDTO(each.ID(), each.Name(), superArtistID, each.Description(), each.FoundedAt(), each.TerminatedAt())
		artists = append(artists, newArtist)
	}

	return context.JSON(http.StatusOK, artists)
}

func NewArtistHandlers(service primary.ArtistManager) *ArtistHandlers {
	return &ArtistHandlers{service: service}
}
