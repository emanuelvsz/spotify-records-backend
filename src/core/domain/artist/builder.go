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
	if superArtistID == uuid.Nil {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.ArtistSuperArtistID,
			Description: messages.ArtistSuperArtistIDInvalidErrMsg,
		})
	}
	b.superArtistID = &superArtistID
	return b
}

func (b *Builder) WithDescription(description string) *Builder {
	if description == "" {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.ArtistDescription,
			Description: messages.ArtistDescriptionInvalidErrMsg,
		})
	}
	b.description = &description
	return b
}

func (b *Builder) WithFoundedAt(foundedAt time.Time) *Builder {
	if foundedAt.IsZero() {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.ArtistFoundedAt,
			Description: messages.ArtistFoundedAtInvalidErrMsg,
		})
	}
	b.foundedAt = foundedAt
	return b
}

func (b *Builder) WithTerminatedAt(terminatedAt *time.Time) *Builder {
	if terminatedAt != nil && terminatedAt.Before(b.foundedAt) {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.ArtistTerminatedAt,
			Description: messages.ArtistTerminatedAtInvalidErrMsg,
		})
	}
	b.terminatedAt = terminatedAt
	return b
}

func (b *Builder) WithSubArtists(subArtists []Artist) *Builder {
	b.subArtists = subArtists
	return b
}

func (b *Builder) WithImageURL(imageURL string) *Builder {
	if imageURL == "" {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.ArtistImageURL,
			Description: messages.ArtistImageURLInvalidErrMsg,
		})
	}
	b.imageURL = &imageURL
	return b
}

func (b *Builder) WithRecordCompanyID(recordCompanyID uuid.UUID) *Builder {
	if recordCompanyID == uuid.Nil {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.ArtistRecordCompanyID,
			Description: messages.ArtistRecordCompanyIDInvalidErrMsg,
		})
	}
	b.recordCompanyID = &recordCompanyID
	return b
}

func (b *Builder) WithCountryID(countryID uuid.UUID) *Builder {
	if countryID == uuid.Nil {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.ArtistCountryID,
			Description: messages.ArtistCountryIDInvalidErrMsg,
		})
	}
	b.countryID = &countryID
	return b
}

func (b *Builder) WithSpotifyURL(spotifyURL string) *Builder {
	if spotifyURL == "" {
		b.invalidFields = append(b.invalidFields, errors.InvalidField{
			Name:        messages.ArtistSpotifyURL,
			Description: messages.ArtistSpotifyURLInvalidErrMsg,
		})
	}
	b.spotifyURL = spotifyURL
	return b
}

func (b *Builder) Build() (*Artist, errors.Error) {
	if len(b.invalidFields) > 0 {
		return nil, errors.NewValidationError(messages.ArtistBuildErr, b.invalidFields...)
	}

	return &b.Artist, nil
}

func NewBuilder() *Builder {
	return &Builder{}
}
