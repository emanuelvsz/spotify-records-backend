package postgres

import (
	"context"
	a "module/src/core/domain/artist"
	s "module/src/core/domain/song"
	"module/src/core/errors"
	"module/src/core/interfaces/repository"
	"module/src/core/messages"
	"module/src/infra/postgres/bridge"

	"github.com/google/uuid"
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
		artistBuilder.WithID(each.ID).WithName(each.Name)
		artistBuilder.WithDescription(each.Description.String).WithFoundedAt(each.FoundedAt)
		artistBuilder.WithTerminatedAt(&each.TerminatedAt.Time)

		var superArtistID uuid.UUID
		if each.SuperArtistID.Valid {
			superArtistID = each.SuperArtistID.UUID
			artistBuilder.WithSuperArtistID(superArtistID)
		} else {
			artistBuilder.WithSuperArtistID(uuid.Nil)
		}

		superID := uuid.NullUUID{
			UUID:  each.ID,
			Valid: true,
		}

		subArtistRows, fetchSubArtistErr := query.SelectSubArtists(context.Background(), superID)
		if fetchSubArtistErr != nil {
			return nil, errors.NewUnexpectedError(messages.FetchingDataErrorMessage, fetchSubArtistErr)
		}

		var subArtist []a.Artist
		for _, j := range subArtistRows {
			subArtistBuilder := a.NewBuilder()
			subArtistBuilder.WithID(j.ID).WithName(j.Name).WithDescription(j.Description.String)
			subArtistBuilder.WithFoundedAt(j.FoundedAt).WithTerminatedAt(&j.TerminatedAt.Time)

			newSubArtist, createSubArtistErr := subArtistBuilder.Build()
			if createSubArtistErr != nil {
				return nil, errors.NewUnexpectedError(messages.FetchingDataErrorMessage, createSubArtistErr)
			}

			subArtist = append(subArtist, *newSubArtist)
		}

		artistBuilder.WithSubArtists(subArtist)
		newArtist, createErr := artistBuilder.Build()
		if createErr != nil {
			return nil, errors.NewUnexpectedError(messages.FetchingDataErrorMessage, createErr)
		}

		artists = append(artists, *newArtist)
	}

	return artists, nil
}

func (ap *ArtistPostgresRepository) FindArtistSongs(artistID uuid.UUID) ([]s.Song, errors.Error) {
	conn, err := ap.getConnection()
	if err != nil {
		return nil, errors.NewUnexpectedError(messages.DataSourceUnavailableErrorMessage, err)
	}
	defer ap.closeConnection(conn)

	query := bridge.New(conn)
	songsRow, fetchErr := query.SelectArtistSongs(context.Background(), artistID)
	if fetchErr != nil {
		return nil, errors.NewUnexpectedError(messages.FetchingDataErrorMessage, fetchErr)
	}

	artists := make([]s.Song, 0)
	for _, each := range songsRow {
		songBuilder := s.NewBuilder()
		songBuilder.WithID(each.ID).WithName(each.Name).WithAlbumID(&each.AlbumID.UUID)
		songBuilder.WithReleaseDate(each.ReleaseDate).WithDuration(each.Duration).WithArtistID(each.ArtistID)

		newSong, createErr := songBuilder.Build()
		if createErr != nil {
			return nil, errors.NewUnexpectedError(messages.FetchingDataErrorMessage, createErr)
		}

		artists = append(artists, *newSong)
	}

	return artists, nil
}

func (ap *ArtistPostgresRepository) FindArtistInformation(artistID uuid.UUID) (*a.Artist, errors.Error) {
	conn, err := ap.getConnection()
	if err != nil {
		return nil, errors.NewUnexpectedError(messages.DataSourceUnavailableErrorMessage, err)
	}
	defer ap.closeConnection(conn)

	query := bridge.New(conn)
	artistRow, fetchErr := query.SelectArtistByID(context.Background(), artistID)
	if fetchErr != nil {
		return nil, errors.NewUnexpectedError(messages.FetchingDataErrorMessage, fetchErr)
	}

	artistBuilder := a.NewBuilder()
	artistBuilder.WithID(artistRow.ID).WithName(artistRow.Name)
	newArtist, createErr := artistBuilder.Build()
	if createErr != nil {
		return nil, errors.NewUnexpectedError(messages.FetchingDataErrorMessage, createErr)
	}

	return newArtist, nil
}

func NewArtistPostgresRepository(manager connectorManager) *ArtistPostgresRepository {
	return &ArtistPostgresRepository{manager}
}
