package album

import (
	"module/src/core/errors"
	"module/src/core/messages"
	"time"

	"github.com/google/uuid"
)

type Builder struct {
	Album
	invalidFields []errors.InvalidField
}

func (b *Builder) WithID(id uuid.UUID) *Builder {
	if id == uuid.Nil {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.AlbumID,
			Description: messages.AlbumIDInvalidErrMsg,
		})
	}
	b.id = id
	return b
}

func (b *Builder) WithName(name string) *Builder {
	if len(name) == 0 {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.AlbumName,
			Description: messages.AlbumNameInvalidErrMsg,
		})
	}
	b.name = name
	return b
}

func (b *Builder) WithArtistID(artistID uuid.UUID) *Builder {
	if artistID == uuid.Nil {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.AlbumArtistID,
			Description: messages.AlbumArtistIDInvalidErrMsg,
		})
	}
	b.artistID = artistID
	return b
}

func (b *Builder) WithReleaseDate(releaseDate time.Time) *Builder {
	if releaseDate.IsZero() {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.AlbumReleaseDate,
			Description: messages.AlbumReleaseDateInvalidErrMsg,
		})
	}
	b.releaseDate = releaseDate
	return b
}

func (b *Builder) Build() (*Album, errors.Error) {
	if len(b.invalidFields) > 0 {
		return nil, errors.NewValidationError(messages.AlbumBuildErr, b.invalidFields...)
	}

	return &b.Album, nil
}

func NewBuilder() *Builder {
	return &Builder{}
}
