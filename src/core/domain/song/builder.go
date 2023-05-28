package song

import (
	"module/src/core/errors"
	"module/src/core/messages"
	"time"

	"github.com/google/uuid"
)

type Builder struct {
	Song
	invalidFields []errors.InvalidField
}

func (b *Builder) WithID(id uuid.UUID) *Builder {
	if id == uuid.Nil {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.SongID,
			Description: messages.SongIDInvalidErrMsg,
		})
	} else {
		b.id = id
	}
	return b
}

func (b *Builder) WithName(name string) *Builder {
	if name == "" {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.SongName,
			Description: messages.SongNameInvalidErrMsg,
		})
	} else {
		b.name = name
	}
	return b
}

func (b *Builder) WithArtistID(artistID uuid.UUID) *Builder {
	if artistID == uuid.Nil {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.SongArtistID,
			Description: messages.SongArtistIDInvalidErrMsg,
		})
	} else {
		b.artistID = artistID
	}
	return b
}

func (b *Builder) WithAlbumID(albumID *uuid.UUID) *Builder {
	b.albumID = albumID
	return b
}

func (b *Builder) WithReleaseDate(releaseDate time.Time) *Builder {
	if releaseDate.IsZero() {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.SongReleaseDate,
			Description: messages.SongReleaseDateInvalidErrMsg,
		})
	} else {
		b.releaseDate = releaseDate
	}
	return b
}

func (b *Builder) WithDuration(duration float32) *Builder {
	if duration <= 0 {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.SongDuration,
			Description: messages.SongDurationInvalidErrMsg,
		})
	} else {
		b.duration = duration
	}
	return b
}

func (b *Builder) Build() (*Song, errors.Error) {
	if len(b.invalidFields) > 0 {
		return nil, errors.NewValidationError(messages.SongBuildErr, b.invalidFields...)
	}

	return &b.Song, nil
}

func NewBuilder() *Builder {
	return &Builder{}
}
