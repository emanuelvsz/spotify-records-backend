package dicontainer

import "module/src/app/api/handlers"

func GetArtistHandlers() *handlers.ArtistHandlers {
	return handlers.NewArtistHandlers(GetArtistHandlers())
}