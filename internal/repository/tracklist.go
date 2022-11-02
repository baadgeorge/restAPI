package repository

import (
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"test/internal/datastruct"
)

type ITracklistQuery interface {
	GetUserTracklist(userId int) ([]datastruct.Track, error)
	AddTrackToUserTracklist(userId, trackId int) error
	DeleteTrackFromUserTracklist(userId, trackId int) error
	DeleteUserTracklist(userId int) error
}

type TracklistQuery struct {
	db *pgx.Conn
}

func NewTracklistQuery(db *pgx.Conn) *TracklistQuery {
	return &TracklistQuery{db: db}
}

func (tl *TracklistQuery) GetUserTracklist(userId int) ([]datastruct.Track, error) {
	var tracksList []datastruct.Track
	rows, err := tl.db.Query("select tracks.* from tracks "+
		"inner join tracklist on tracklist.track_id = tracks.track_id where tracklist.user_id = $1", userId)
	for rows.Next() {
		var track datastruct.Track
		err := rows.Scan(&track.ID, &track.Author, &track.Title, &track.Genre, &track.Album, &track.Duration)
		if err != nil {
			logrus.Debug(err)
			return nil, err
		}
		tracksList = append(tracksList, track)
	}

	return tracksList, err
}

func (tl *TracklistQuery) AddTrackToUserTracklist(userId, trackId int) error {
	_, err := tl.db.Exec("insert into tracklist (track_id, user_id) values ($1, $2)", trackId, userId)
	if err != nil {
		logrus.Debug(err, userId, trackId)
		return err
	}
	return nil
}

func (tl *TracklistQuery) DeleteTrackFromUserTracklist(userId, trackId int) error {
	_, err := tl.db.Exec("delete from tracklist where user_id = $1 and track_id = $2", userId, trackId)
	if err != nil {
		logrus.Debug(err, userId, trackId)
		return err
	}
	return nil
}

func (tl *TracklistQuery) DeleteUserTracklist(userId int) error {
	_, err := tl.db.Exec("delete from tracklist where user_id = $1", userId)
	if err != nil {
		logrus.Debug(err, userId)
		return err
	}
	return nil
}
