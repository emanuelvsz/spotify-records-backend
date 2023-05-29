package routes

import (
	"module/src/app/api/endpoints/dicontainer"

	"github.com/labstack/echo/v4"
)

func loadArtistRoutes(group *echo.Group) {
	artistHandlers := dicontainer.GetArtistHandlers()

	group.GET("/artists", artistHandlers.GetArtists)
	group.GET("/artists/:artistID/songs", artistHandlers.GetArtistSongs)
}
