package postgres

import (
	"context"
	"fmt"
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

	for _, each := range artistRows {
		fmt.Println(each)
	}

	return nil, nil
}

func NewArtistPostgresRepository(manager connectorManager) *ArtistPostgresRepository {
	return &ArtistPostgresRepository{manager}
}
