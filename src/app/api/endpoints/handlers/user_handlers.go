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

const (
	albumID  = "albumID"
	artistID = "artistID"
	songID   = "songID"
)

type UserHandlers struct {
	service primary.UserManager
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
// @Router /user/artists [get]
func (h UserHandlers) GetArtists(context echo.Context) error {
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
// @Router /user/artists/{artistID}/songs [get]
func (h UserHandlers) GetArtistSongs(context echo.Context) error {
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
// @Param artistID path string true "ID do artista." default(a6d6488b-6ed0-4d41-9c55-4e899fdd7e47)
// @Produce json
// @Success 200 {array} response.ArtistLowDTO "Requisição realizada com sucesso."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 422 {object} response.ErrorMessage "Algum dado informado não pôde ser processado."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado."
// @Failure 503 {object} response.ErrorMessage "A base de dados não está disponível."
// @Router /user/artists/{artistID} [get]
func (h UserHandlers) GetArtistInformation(context echo.Context) error {
	artistIDParam, conversionErr := converters.ConvertFromStringToUUID(context.Param(artistID), messages.ArtistID,
		messages.ArtistIDInvalidErrMsg, messages.ConversionErrorMessage)
	if conversionErr != nil {
		return getHttpHandledErrorResponse(context, conversionErr)
	}

	artist, fetchErr := h.service.FetchArtistInformation(*artistIDParam)
	if fetchErr != nil {
		return getHttpHandledErrorResponse(context, fetchErr)
	}

	subArtists := make([]response.ArtistDTO, 0)
	for _, each := range artist.SubArtists() {
		artistDTO := response.NewArtistDTO(
			each.ID(),
			each.Name(),
			each.SuperArtistID(),
			each.Description(),
			each.FoundedAt(),
			each.TerminatedAt(),
			nil,
		)

		subArtists = append(subArtists, *artistDTO)
	}

	recordComID := artist.RecordCompanyID()
	if *recordComID == uuid.Nil {
		recordComID = nil
	}

	countryID := artist.CountryID()
	if *countryID == uuid.Nil {
		countryID = nil
	}

	artistDTO := response.NewArtistLowDTO(
		artist.ID(),
		artist.Name(),
		artist.SuperArtistID(),
		artist.Description(),
		artist.FoundedAt(),
		artist.TerminatedAt(),
		subArtists,
		artist.ImageURL(),
		recordComID,
		countryID,
		*artist.SpotifyURL(),
	)

	return context.JSON(http.StatusOK, artistDTO)
}

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
// @Router /user/albums/{albumID}/songs [get]
func (h UserHandlers) GetAlbumSongs(context echo.Context) error {
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
// @Router /user/albums [get]
func (h UserHandlers) GetAlbums(context echo.Context) error {
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

func NewUserHandlers(service primary.UserManager) *UserHandlers {
	return &UserHandlers{service: service}
}
