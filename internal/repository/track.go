package repository

import (
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"test/internal/datastruct"
)

type ITrackQuery interface {
	CreateTrack(track datastruct.Track) (int, error)
	DeleteTrack(id int) error
	UpdateTrack(track datastruct.Track) error
	GetTrackById(id int) (*datastruct.Track, error)
	GetTracksByAuthorAndTitle(author string, title string) ([]datastruct.Track, error)
	GetTracksByAuthor(author string) ([]datastruct.Track, error)
	GetTracksByTitle(title string) ([]datastruct.Track, error)
}

type TrackQuery struct {
	db *pgx.Conn
}

func NewTrackQuery(db *pgx.Conn) *TrackQuery {
	return &TrackQuery{db: db}
}

func (t *TrackQuery) CreateTrack(track datastruct.Track) (int, error) {
	var trackId int
	err := t.db.QueryRow("insert into tracks (author, title, genre, album, duration) "+
		"values($1, $2, $3, $4, $5) returning track_id",
		track.Author, track.Title, track.Genre, track.Album, track.Duration).Scan(&trackId)
	if err != nil {
		logrus.Debug(err)
		return 0, err
	}
	return trackId, nil
}

func (t *TrackQuery) DeleteTrack(trackId int) error {
	_, err := t.db.Exec("delete from tracks where track_id=$1", trackId)
	if err != nil {
		logrus.Debug(err)
		return err
	}
	return nil

}
func (t *TrackQuery) UpdateTrack(track datastruct.Track) error {
	rows, err := t.db.Query("update tracks set author=$1, title=$2, genre=$3, album=$4, duration=$5 "+
		"where id=$6 returning track_id, author, title, genre, album, duration",
		track.Author, track.Title, track.Genre, track.Album, track.Duration, track.ID)
	if err != nil {
		logrus.Debug(err)
		return err
	}
	var updTrack datastruct.Track
	err = rows.Scan(&updTrack.ID, &updTrack.Author, &updTrack.Title, &updTrack.Genre, &updTrack.Album, &updTrack.Duration)
	if err != nil {
		logrus.Debug(err)
		return err
	}
	return nil
}

func (t *TrackQuery) GetTrackById(trackId int) (*datastruct.Track, error) {
	var track datastruct.Track
	row := t.db.QueryRow("select * from tracks where track_id=$1", trackId)
	err := row.Scan(&track.ID, &track.Author, &track.Title, &track.Album, &track.Genre, &track.Duration)
	if err != nil {
		logrus.Debug(err)
		return nil, err
	}
	return &track, nil
}

func (t *TrackQuery) GetTracksByAuthorAndTitle(author string, title string) ([]datastruct.Track, error) {
	rows, err := t.db.Query("select * from tracks where author=$1 and title=$2", author, title)
	if err != nil {
		logrus.Debug(err)
		return nil, err
	}
	var trackSlice []datastruct.Track
	for rows.Next() {
		var curTrack datastruct.Track
		err := rows.Scan(&curTrack.ID, &curTrack.Author, &curTrack.Title, &curTrack.Genre, &curTrack.Album, &curTrack.Duration)
		if err != nil {
			logrus.Debug(err)
			return nil, err
		}
		trackSlice = append(trackSlice, curTrack)
	}
	return trackSlice, nil
}
func (t *TrackQuery) GetTracksByAuthor(author string) ([]datastruct.Track, error) {
	rows, err := t.db.Query("select * from tracks where author=$1", author)
	if err != nil {
		logrus.Debug(err)
		return nil, err
	}
	var trackSlice []datastruct.Track
	for rows.Next() {
		var curTrack datastruct.Track
		err := rows.Scan(&curTrack.ID, &curTrack.Author, &curTrack.Title, &curTrack.Genre, &curTrack.Album, &curTrack.Duration)
		if err != nil {
			logrus.Debug(err)
			return nil, err
		}
		trackSlice = append(trackSlice, curTrack)
	}
	return trackSlice, nil

}
func (t *TrackQuery) GetTracksByTitle(title string) ([]datastruct.Track, error) {
	rows, err := t.db.Query("select * from tracks where title=$1", title)
	if err != nil {
		return nil, err
	}
	var trackSlice []datastruct.Track
	for rows.Next() {
		var curTrack datastruct.Track
		err := rows.Scan(&curTrack.ID, &curTrack.Author, &curTrack.Title, &curTrack.Genre, &curTrack.Album, &curTrack.Duration)
		if err != nil {
			return nil, err
		}
		trackSlice = append(trackSlice, curTrack)
	}
	return trackSlice, nil
}
