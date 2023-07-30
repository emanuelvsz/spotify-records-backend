package dicontainer

import (
	"module/src/core/interfaces/repository"
	"module/src/infra/postgres"
)

func GetAuthRepository() repository.AuthLoader {
	return postgres.NewAuthPostgresRepository(GetPsqlConnectionManager())
}

func GetUserRepository() repository.UserLoader {
	return postgres.NewUserPostgresRepository(GetPsqlConnectionManager())
}

func GetPsqlConnectionManager() *postgres.DatabaseConnectionManager {
	return &postgres.DatabaseConnectionManager{}
}
