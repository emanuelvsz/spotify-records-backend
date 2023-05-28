package postgres

import (
	u "module/src/core/domain/user"
	"module/src/core/domain/user/credentials"
	"module/src/core/errors"
	"module/src/core/interfaces/repository"
	"module/src/core/messages"
	"strings"
)

var _ repository.AuthLoader = &AuthPostgresRepository{}

type AuthPostgresRepository struct {
	connectorManager
}

func (instance *AuthPostgresRepository) Login(credentials credentials.Credentials) (*u.User, errors.Error) {
	return nil, nil
}

func (instance AuthPostgresRepository) handleError(err error) errors.Error {
	msg := err.Error()

	if strings.Contains(msg, "sql: no rows in result set") {
		return errors.NewNotFoundError(messages.InvalidCredentialsErrorMessage, err)
	}

	return errors.NewUnexpectedError(messages.UnexpectedErrorMessage, err)
}

func NewAuthPostgresRepository(manager connectorManager) *AuthPostgresRepository {
	return &AuthPostgresRepository{manager}
}
