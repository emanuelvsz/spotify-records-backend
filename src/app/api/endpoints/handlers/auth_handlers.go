package handlers

import (
	"module/src/core/interfaces/primary"
)

type AuthHandlers struct {
	service primary.AuthManager
}

func NewAuthHandlers(service primary.AuthManager) *AuthHandlers {
	return &AuthHandlers{service: service}
}
