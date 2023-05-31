package handlers

import (
	"module/src/app/api/endpoints/handlers/converters"
	"module/src/app/api/endpoints/handlers/dtos/response"
	"module/src/core/interfaces/primary"
	"module/src/core/messages"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AlbumHandlers struct {
	service primary.AlbumManager
}

const (
	albumID = "albumID"
)

// GetAlbumSongs
// @ID GetAlbumSongs
// @Summary Buscar todas as músicas de um album específico
// @Tags Rotas do usuário
// @Description Rota que permite que se busque todas as músicas de um album
// @Param albumID path string true "ID do album." default(4d1035ba-22b7-4139-baff-c02861e4c6ec)
// @Produce json
// @Success 200 {array} response.SongDTO "Requisição realizada com sucesso."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 422 {object} response.ErrorMessage "Algum dado informado não pôde ser processado."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado."
// @Failure 503 {object} response.ErrorMessage "A base de dados não está disponível."
// @Router /albums/{albumID}/songs [get]
func (h AlbumHandlers) GetAlbumSongs(context echo.Context) error {
	albumID, conversionErr := converters.ConvertFromStringToUUID(context.Param(albumID), messages.AlbumID,
		messages.AlbumIDInvalidErrMsg, messages.ConversionErrorMessage)
	if conversionErr != nil {
		return getHttpHandledErrorResponse(context, conversionErr)
	}

	songsRow, fetchErr := h.service.FetchAlbumSongs(*albumID)
	if fetchErr != nil {
		return getHttpHandledErrorResponse(context, fetchErr)
	}

	var songs []response.SongHighDTO
	for _, each := range songsRow {
		url := each.SpotifyURL()

		songBuilder := response.NewSongHighDTO(
			each.ID(),
			each.Name(),
			each.AlbumID(),
			each.ReleaseDate(),
			each.Duration(),
			each.Lyrics(),
			&url,
			each.TrackNumber(),
		)
		songs = append(songs, *songBuilder)
	}

	return context.JSON(http.StatusOK, songs)
}

// GetAlbums
// @ID GetAlbums
// @Summary Buscar todos os albuns
// @Tags Rotas do usuário
// @Description Rota que permite que se busque todos os albuns
// @Produce json
// @Success 200 {array} response.AlbumDTO "Requisição realizada com sucesso."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 422 {object} response.ErrorMessage "Algum dado informado não pôde ser processado."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado."
// @Failure 503 {object} response.ErrorMessage "A base de dados não está disponível."
// @Router /albums [get]
func (h AlbumHandlers) GetAlbums(context echo.Context) error {
	albumsRow, fetchErr := h.service.FetchAlbums()
	if fetchErr != nil {
		return getHttpHandledErrorResponse(context, fetchErr)
	}

	var albums []response.AlbumDTO
	for _, each := range albumsRow {
		url := each.ImageURL()
		albumBuilder := response.NewAlbumDTO(
			each.ID(),
			each.Name(),
			each.ArtistID(),
			each.ReleaseDate(),
			each.Description(),
			&url,
		)

		albums = append(albums, *albumBuilder)
	}

	return context.JSON(http.StatusOK, albums)
}

func NewAlbumHandlers(service primary.AlbumManager) *AlbumHandlers {
	return &AlbumHandlers{service: service}
}
