package dicontainer

import "module/src/api/handlers"

func GetArtistHandlers() *handlers.ArtistHandlers {
	return handlers.NewArtistHandlers(GetArtistHandlers())
}