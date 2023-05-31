package dicontainer

import (
	"module/src/core/errors/logger"
	"module/src/core/interfaces/primary"
	"module/src/core/services"
)

func GetAuthServices() primary.AuthManager {
	return services.NewAuthServices(GetAuthRepository(), GetLogger())
}

func GetArtistServices() primary.ArtistManager {
	return services.NewArtistServices(GetArtistRepository(), GetLogger())
}

func GetAlbumServices() primary.AlbumManager {
	return services.NewAlbumServices(GetAlbumRepository(), GetLogger())
}

func GetLogger() logger.Logger {
	return logger.New()
}
