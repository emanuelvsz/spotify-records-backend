package postgres

import (
	"context"
	"module/src/core/domain/album"
	s "module/src/core/domain/song"
	"module/src/core/errors"
	"module/src/core/interfaces/repository"
	"module/src/core/messages"
	"module/src/infra/postgres/bridge"

	"github.com/google/uuid"
)

var _ repository.AlbumLoader = &AlbumPostgresRepository{}

type AlbumPostgresRepository struct {
	connectorManager
}

func (a AlbumPostgresRepository) FindAlbumSongs(albumID uuid.UUID) ([]s.Song, errors.Error) {
	conn, err := a.getConnection()
	if err != nil {
		return nil, errors.NewUnexpectedError(messages.DataSourceUnavailableErrorMessage, err)
	}
	defer a.closeConnection(conn)

	query := bridge.New(conn)

	nullAlbumID := uuid.NullUUID{
		UUID:  albumID,
		Valid: true,
	}

	songsRow, fetchErr := query.SelectAlbumSongs(context.Background(), nullAlbumID)
	if fetchErr != nil {
		return nil, errors.NewUnexpectedError(messages.FetchingDataErrorMessage, fetchErr)
	}

	songs := make([]s.Song, 0)
	for _, each := range songsRow {
		songBuilder := s.NewBuilder()
		songBuilder.WithAlbumID(&each.AlbumID.UUID).WithName(each.Name).WithDuration(each.Duration)
		songBuilder.WithLyrics(each.Lyrics.String).WithReleaseDate(each.ReleaseDate)
		songBuilder.WithTrackNumber(int(each.TrackNumber.Int32)).WithSpotifyURL(each.SpotifyUrl.String)
		songBuilder.WithID(each.ID)

		newSong, createErr := songBuilder.Build()
		if createErr != nil {
			return nil, errors.NewUnexpectedError(messages.FetchingDataErrorMessage, createErr)
		}

		songs = append(songs, *newSong)
	}

	return songs, nil
}

func (a AlbumPostgresRepository) FindAlbums() ([]album.Album, errors.Error) {
	conn, err := a.getConnection()
	if err != nil {
		return nil, errors.NewUnexpectedError(messages.DataSourceUnavailableErrorMessage, err)
	}
	defer a.closeConnection(conn)

	query := bridge.New(conn)
	albumRows, fetchErr := query.SelectAlbums(context.Background())
	if fetchErr != nil {
		return nil, errors.NewUnexpectedError(messages.FetchingDataErrorMessage, fetchErr)
	}

	albums := make([]album.Album, 0)
	for _, each := range albumRows {
		albumBuilder := album.NewBuilder()
		albumBuilder.WithID(each.ID).WithName(each.Name).WithArtistID(each.ArtistID)
		albumBuilder.WithDescription(each.Description.String).WithReleaseDate(each.ReleaseDate)
		albumBuilder.WithImageURL(each.ImageUrl.String)

		newAlbum, createErr := albumBuilder.Build()
		if createErr != nil {
			return nil, errors.NewUnexpectedError(messages.FetchingDataErrorMessage, createErr)
		}

		albums = append(albums, *newAlbum)
	}

	return albums, nil
}

func NewAlbumPostgresRepository(manager connectorManager) *AlbumPostgresRepository {
	return &AlbumPostgresRepository{manager}
}
