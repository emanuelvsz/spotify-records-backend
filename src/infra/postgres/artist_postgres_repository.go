package postgres

import (
	"context"
	a "module/src/core/domain/artist"
	"module/src/core/errors"
	"module/src/core/interfaces/repository"
	"module/src/core/messages"
	"module/src/infra/postgres/bridge"
)

var _ repository.ArtistLoader = &ArtistPostgresRepository{}

type ArtistPostgresRepository struct {
	connectorManager
}

func (ap *ArtistPostgresRepository) FindArtists() ([]a.Artist, errors.Error) {
	conn, err := ap.getConnection()
	if err != nil {
		return nil, errors.NewUnexpectedError(messages.DataSourceUnavailableErrorMessage, err)
	}
	defer ap.closeConnection(conn)

	query := bridge.New(conn)

	artistRows, fetchErr := query.SelectArtists(context.Background())
	if fetchErr != nil {
		return nil, errors.NewUnexpectedError(messages.FetchingDataErrorMessage, fetchErr)
	}

	artists := make([]a.Artist, 0)
	for _, each := range artistRows {
		artistBuilder := a.NewBuilder()
		artistBuilder.WithID(each.ID).WithName(each.Name).WithSuperArtistID(each.SuperArtistID.UUID)
		artistBuilder.WithDescription(each.Description.String).WithFoundedAt(each.FoundedAt).WithTerminatedAt(&each.TerminatedAt.Time)
	
		newArtist, createErr := artistBuilder.Build()
		if createErr != nil {
			return nil, errors.NewUnexpectedError(messages.FetchingDataErrorMessage, createErr)
		}

		artists = append(artists, *newArtist)
	}

	return artists, nil
}

func NewArtistPostgresRepository(manager connectorManager) *ArtistPostgresRepository {
	return &ArtistPostgresRepository{manager}
}
