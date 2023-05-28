package services

import (
	userPkg "module/src/core/domain/user"
	"module/src/core/domain/user/credentials"
	"module/src/core/errors"
	"module/src/core/errors/logger"
	"module/src/core/interfaces/primary"
	"module/src/core/interfaces/repository"

	"github.com/google/uuid"
)

var _ primary.AuthManager = (*AuthServices)(nil)

type AuthServices struct {
	authRepository repository.AuthLoader
	logger         logger.Logger
}

func (a *AuthServices) Login(credentials credentials.Credentials) (*userPkg.User, errors.Error) {
	return nil, nil
}

func (a *AuthServices) SessionExists(userID uuid.UUID, token string) (bool, errors.Error) {
	return true, nil
}

func NewAuthServices(authRepository repository.AuthLoader, logger logger.Logger) *AuthServices {
	return &AuthServices{
		authRepository: authRepository,
		logger:         logger,
	}
}