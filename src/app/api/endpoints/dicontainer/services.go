package dicontainer

import (
	"module/src/core/errors/logger"
	"module/src/core/interfaces/primary"
	"module/src/core/services"
)

func GetAuthServices() primary.AuthManager {
	return services.NewAuthServices(GetAuthRepository(), GetLogger())
}

func GetUserServices() primary.UserManager {
	return services.NewUserServices(GetUserRepository(), GetLogger())
}

func GetLogger() logger.Logger {
	return logger.New()
}
