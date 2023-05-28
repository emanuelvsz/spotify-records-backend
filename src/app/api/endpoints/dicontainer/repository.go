package dicontainer

import (
	"module/src/core/interfaces/repository"
	"module/src/infra/postgres"
)

func GetAuthRepository() repository.AuthLoader {
	return postgres.NewAuthPostgresRepository(GetPsqlConnectionManager())
}

func GetArtistRepository() repository.ArtistLoader {
	return postgres.NewArtistPostgresRepository(GetPsqlConnectionManager())
}


func GetPsqlConnectionManager() *postgres.DatabaseConnectionManager {
	return &postgres.DatabaseConnectionManager{}
}
