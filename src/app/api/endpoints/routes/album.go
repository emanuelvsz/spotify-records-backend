package routes

import (
	"module/src/app/api/endpoints/dicontainer"

	"github.com/labstack/echo/v4"
)

func loadAlbumRoutes(group *echo.Group) {
	albumGroup := group.Group("/albums")
	albumHandlers := dicontainer.GetAlbumHandlers()

	albumGroup.GET("", albumHandlers.GetAlbums)
	albumGroup.GET("/:albumID/songs", albumHandlers.GetAlbumSongs)
}
