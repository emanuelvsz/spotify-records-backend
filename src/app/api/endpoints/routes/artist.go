package routes

import (
	"module/src/app/api/endpoints/dicontainer"

	"github.com/labstack/echo/v4"
)

func loadArtistRoutes(group *echo.Group) {
	artistGroup := group.Group("/artists")
	artistHandlers := dicontainer.GetArtistHandlers()

	artistGroup.GET("", artistHandlers.GetArtists)
	artistGroup.GET("/:artistID", artistHandlers.GetArtistInformation)
	artistGroup.GET("/:artistID/songs", artistHandlers.GetArtistSongs)
}
