package artist

import "github.com/google/uuid"

type Artist struct {
	ID          uuid.UUID
	Name        string
	Age         int8
	Nacionality string
}
