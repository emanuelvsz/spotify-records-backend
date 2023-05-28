package artist

import (
	"module/src/core/errors"
	"module/src/core/messages"
	"time"

	"github.com/google/uuid"
)

type Builder struct {
	Artist
	invalidFields []errors.InvalidField
}

func (b *Builder) WithID(id uuid.UUID) *Builder {
	if id == uuid.Nil {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.ArtistID,
			Description: messages.ArtistIDInvalidErrMsg,
		})
	}
	b.id = id
	return b
}

func (b *Builder) WithName(name string) *Builder {
	if name == "" {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.ArtistName,
			Description: messages.ArtistNameInvalidErrMsg,
		})
	}
	b.name = name
	return b
}

func (b *Builder) WithSuperArtistID(superArtistID uuid.UUID) *Builder {
	b.superArtistID = &superArtistID
	return b
}

func (b *Builder) WithDescription(description string) *Builder {
	b.description = &description
	return b
}

func (b *Builder) WithFoundedAt(foundedAt time.Time) *Builder {
	b.foundedAt = foundedAt
	return b
}

func (b *Builder) WithTerminatedAt(terminatedAt *time.Time) *Builder {
	b.terminatedAt = terminatedAt
	return b
}

func (b *Builder) validateDates() {
	if b.terminatedAt != nil && b.foundedAt.After(*b.terminatedAt) {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.ArtistFoundedAt,
			Description: messages.ArtistFoundedAtErrMsg,
		})
	}
}

func (b *Builder) Build() (*Artist, errors.Error) {
	b.validateDates()
	if len(b.invalidFields) > 0 {
		return nil, errors.NewValidationError(messages.ArtistBuildErr, b.invalidFields...)
	}

	return &b.Artist, nil
}

func NewBuilder() *Builder {
	return &Builder{}
}
