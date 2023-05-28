package repository

import (
	"module/src/core/domain/user"
	"module/src/core/domain/user/credentials"
	"module/src/core/errors"
)

type AuthLoader interface {
	Login(credentials credentials.Credentials) (*user.User, errors.Error)
}
