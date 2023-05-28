package primary

import (
	"module/src/core/domain/user"
	"module/src/core/domain/user/credentials"
	"module/src/core/errors"

	"github.com/google/uuid"
)

type AuthManager interface {
	Login(credentials credentials.Credentials) (*user.User, errors.Error)
	SessionExists(userID uuid.UUID, token string) (bool, errors.Error)
}
