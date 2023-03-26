package song

import "github.com/google/uuid"

type Song struct {
	ID   uuid.UUID
	Name string
}
