package dicontainer

import "module/src/app/api/endpoints/handlers"

func GetAuthHandlers() *handlers.AuthHandlers {
	return handlers.NewAuthHandlers(GetAuthServices())
}

func GetArtistHandlers() *handlers.ArtistHandlers {
	return handlers.NewArtistHandlers(GetArtistServices())
}