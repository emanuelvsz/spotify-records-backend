package album

import (
	"module/src/core/domain"
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
	if name == "" {
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

func (b *Builder) WithDescription(description string) *Builder {
	if description == "" {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.AlbumDescription,
			Description: messages.AlbumDescriptionInvalidErrMsg,
		})
	}
	b.description = &description
	return b
}

func (b *Builder) WithImageURL(imageURL string) *Builder {
	if imageURL == "" {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.AlbumImageURL,
			Description: messages.AlbumImageURLInvalidErrMsg,
		})
	}
	b.imageURL = imageURL
	return b
}

func (b *Builder) Build() (*Album, errors.Error) {
	if len(b.invalidFields) > 0 {
		domain.ShowInvalidFields(b.invalidFields)
		return nil, errors.NewValidationError(messages.AlbumBuildErr, b.invalidFields...)
	}

	return &b.Album, nil
}

func NewBuilder() *Builder {
	return &Builder{}
}
