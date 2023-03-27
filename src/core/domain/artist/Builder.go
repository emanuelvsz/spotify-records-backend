package artist

import (
	"module/src/core/errors"
	"module/src/core/messages"
	"strings"
)

type builder struct {
	artist        *Artist
	invalidFields []errors.InvalidFields
}

func NewBuilder() *builder {
	return &builder{artist: &Artist{}}
}

func (instance *builder) WithName(name string) *builder {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		instance.invalidFields = append(instance.invalidFields, errors.InvalidFields{
			Name:        messages.ArtistName,
			Description: messages.InvalidArtistName,
		})
	}

	return instance
}

func (instance *builder) Build() (*Artist, error) {
	if len(instance.invalidFields) > 0 {
		return nil, nil
	}

	return instance.artist, nil
}
