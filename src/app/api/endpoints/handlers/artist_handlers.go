package handlers

import (
	"module/src/app/api/endpoints/handlers/converters"
	"module/src/app/api/endpoints/handlers/dtos/response"
	"module/src/core/interfaces/primary"
	"module/src/core/messages"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ArtistHandlers struct {
	service primary.ArtistManager
}

const (
	artistID = "artistID"
)

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

		subArtists := make([]response.ArtistDTO, 0)
		for _, j := range each.SubArtists() {
			newSubArtist := *response.NewArtistDTO(j.ID(), j.Name(), j.SuperArtistID(), j.Description(), j.FoundedAt(), each.TerminatedAt(), nil)
			subArtists = append(subArtists, newSubArtist)
		}

		newArtist := *response.NewArtistDTO(each.ID(), each.Name(), superArtistID, each.Description(), each.FoundedAt(), each.TerminatedAt(), subArtists)
		artists = append(artists, newArtist)
	}

	return context.JSON(http.StatusOK, artists)
}

// GetArtistSongs
// @ID GetArtistSongs
// @Summary Buscar todas as músicas de um artista específico
// @Tags Rotas do usuário
// @Description Rota que permite que se busque todas as músicas de um artista
// @Param artistID path string true "ID do artista." default(cfd8b073-a303-4886-836a-f249e88be9bc)
// @Produce json
// @Success 200 {array} response.SongDTO "Requisição realizada com sucesso."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 422 {object} response.ErrorMessage "Algum dado informado não pôde ser processado."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado."
// @Failure 503 {object} response.ErrorMessage "A base de dados não está disponível."
// @Router /artists/{artistID}/songs [get]
func (h ArtistHandlers) GetArtistSongs(context echo.Context) error {
	artistID, conversionErr := converters.ConvertFromStringToUUID(context.Param(artistID), messages.ArtistID,
		messages.ArtistIDInvalidErrMsg, messages.ConversionErrorMessage)
	if conversionErr != nil {
		return getHttpHandledErrorResponse(context, conversionErr)
	}

	songsRow, fetchErr := h.service.FetchArtistSongs(*artistID)
	if fetchErr != nil {
		return getHttpHandledErrorResponse(context, fetchErr)
	}

	var songs []response.SongDTO
	for _, each := range songsRow {
		songBuilder := response.NewSongDTO(each.ID(), each.Name(), each.ArtistID(), each.AlbumID(), each.ReleaseDate(), each.Duration())
		songs = append(songs, *songBuilder)
	}

	return context.JSON(http.StatusOK, songs)
}

// GetArtistInformation
// @ID GetArtistInformation
// @Summary Buscar os dados de um artista por id
// @Tags Rotas do usuário
// @Description Rota que permite que se busque as informações de um artista
// @Param artistID path string true "ID do artista." default(cfd8b073-a303-4886-836a-f249e88be9bc)
// @Produce json
// @Success 200 {array} response.ArtistDTO "Requisição realizada com sucesso."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 422 {object} response.ErrorMessage "Algum dado informado não pôde ser processado."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado."
// @Failure 503 {object} response.ErrorMessage "A base de dados não está disponível."
// @Router /artists/{artistID} [get]
func (h ArtistHandlers) GetArtistInformation(context echo.Context) error {
	artistID, conversionErr := converters.ConvertFromStringToUUID(context.Param(artistID), messages.ArtistID,
		messages.ArtistIDInvalidErrMsg, messages.ConversionErrorMessage)
	if conversionErr != nil {
		return getHttpHandledErrorResponse(context, conversionErr)
	}

	artist, fetchErr := h.service.FetchArtistInformation(*artistID)
	if fetchErr != nil {
		return getHttpHandledErrorResponse(context, fetchErr)
	}
	artistDTO := response.NewArtistDTO(artist.ID(), artist.Name(), artist.SuperArtistID(), artist.Description(), artist.FoundedAt(), artist.TerminatedAt(), nil)

	return context.JSON(http.StatusOK, artistDTO)
}

func NewArtistHandlers(service primary.ArtistManager) *ArtistHandlers {
	return &ArtistHandlers{service: service}
}
