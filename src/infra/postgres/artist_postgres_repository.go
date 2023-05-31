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

		var superArtistID uuid.UUID
		if each.SuperArtistID.Valid {
			superArtistID = each.SuperArtistID.UUID
			artistBuilder.WithSuperArtistID(&superArtistID)
		} else {
			artistBuilder.WithSuperArtistID(&uuid.Nil)
		}

		subArtistRows, fetchSubArtistErr := query.SelectSubArtists(context.Background(), each.ID)
		if fetchSubArtistErr != nil {
			return nil, errors.NewUnexpectedError(messages.FetchingDataErrorMessage, fetchSubArtistErr)
		}

		subArtist := make([]a.Artist, 0)
		for _, j := range subArtistRows {
			subArtistBuilder := a.NewBuilder()
			subArtistBuilder.WithID(j.ID).WithName(j.Name)

			subDescription := j.Description.String
			subArtistBuilder.WithDescription(&subDescription)

			subArtistBuilder.WithFoundedAt(j.FoundedAt).WithTerminatedAt(&j.TerminatedAt.Time)

			newSubArtist, createSubArtistErr := subArtistBuilder.Build()
			if createSubArtistErr != nil {
				return nil, errors.NewUnexpectedError(messages.FetchingDataErrorMessage, createSubArtistErr)
			}

			subArtist = append(subArtist, *newSubArtist)
		}

		artistBuilder.WithID(each.ID).WithName(each.Name)

		description := each.Description.String
		artistBuilder.WithDescription(&description)

		artistBuilder.WithFoundedAt(each.FoundedAt).WithTerminatedAt(&each.TerminatedAt.Time)
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
	row, fetchErr := query.SelectArtistByID(context.Background(), artistID)
	if fetchErr != nil {
		return nil, errors.NewUnexpectedError(messages.FetchingDataErrorMessage, fetchErr)
	}

	artistBuilder := a.NewBuilder()
	artistBuilder.WithID(row.ID).WithName(row.Name).WithDescription(&row.Description.String)
	artistBuilder.WithCountryID(&row.CountryID.UUID).WithImageURL(&row.ImageUrl.String).WithFoundedAt(row.FoundedAt)
	artistBuilder.WithTerminatedAt(&row.TerminatedAt.Time).WithRecordCompanyID(&row.RecordCompanyID.UUID)
	artistBuilder.WithSpotifyURL(&row.SpotifyUrl.String)

	subArtistRows, fetchSubArtistsErr := query.SelectSubArtists(context.Background(), artistID)
	if fetchSubArtistsErr != nil {
		return nil, errors.NewUnexpectedError(messages.FetchingDataErrorMessage, fetchSubArtistsErr)
	}

	subArtist := make([]a.Artist, 0)

	if len(subArtistRows) != 0 {
		for _, j := range subArtistRows {
			subArtistBuilder := a.NewBuilder()
			subArtistBuilder.WithID(j.ID).WithName(j.Name)

			subDescription := j.Description.String
			subArtistBuilder.WithDescription(&subDescription)

			subArtistBuilder.WithFoundedAt(j.FoundedAt).WithTerminatedAt(&j.TerminatedAt.Time)

			newSubArtist, createSubArtistErr := subArtistBuilder.Build()
			if createSubArtistErr != nil {
				return nil, errors.NewUnexpectedError(messages.FetchingDataErrorMessage, createSubArtistErr)
			}

			subArtist = append(subArtist, *newSubArtist)
		}
	}

	artistBuilder.WithSubArtists(subArtist)

	//TODO: Error here. Probably are some logic that is not working

	newArtist, createErr := artistBuilder.Build()
	if createErr != nil {
		return nil, errors.NewUnexpectedError(messages.FetchingDataErrorMessage, createErr)
	}

	return newArtist, nil
}

func NewArtistPostgresRepository(manager connectorManager) *ArtistPostgresRepository {
	return &ArtistPostgresRepository{manager}
}
