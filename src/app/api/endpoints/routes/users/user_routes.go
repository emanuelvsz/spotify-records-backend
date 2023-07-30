package users

import (
	"module/src/app/api/endpoints/dicontainer"

	"github.com/labstack/echo/v4"
)

var userHandlers = dicontainer.GetUserHandlers() 

func LoadUserRoutes(group *echo.Group) {
	userGroup := group.Group("/user")

	loadUserAlbumRoutes(userGroup)
	loadUserArtistRoutes(userGroup)
}

func loadUserAlbumRoutes(group *echo.Group) {
	albumGroup := group.Group("/album")

	albumGroup.GET("", userHandlers.GetAlbums)
	albumGroup.GET("/:albumID/songs", userHandlers.GetAlbumSongs)
}

func loadUserArtistRoutes(group *echo.Group) {
	artistGroup := group.Group("/artist")

	artistGroup.GET("", userHandlers.GetArtists)
	artistGroup.GET("/:artistID", userHandlers.GetArtistInformation)
	artistGroup.GET("/:artistID/songs", userHandlers.GetArtistSongs)
}
