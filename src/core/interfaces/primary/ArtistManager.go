package primary

import "github.com/labstack/echo/v4"

type ArtistManager interface {
	ListArtists(context echo.Context) error
}
