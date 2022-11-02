package service

import (
	"github.com/sirupsen/logrus"
	"test/internal/datastruct"
	"test/internal/repository"
)

type ITracklistService interface {
	GetUserTracklist(userId int) ([]datastruct.Track, error)
	AddTrackToUserTrackList(userId, trackId int) error
	DeleteTrackFromUserTracklist(userId, trackId int) error
	DeleteUserTrackList(userId int) error
}

type TracklistService struct {
	repo repository.ITracklistQuery
}

func NewTrackListService(repo repository.ITracklistQuery) ITracklistService {
	return &TracklistService{repo: repo}
}

func (ls *TracklistService) GetUserTracklist(userId int) ([]datastruct.Track, error) {
	tracklist, err := ls.repo.GetUserTracklist(userId)
	if err != nil {
		logrus.Debug(err)
		return nil, err
	}
	return tracklist, nil
}

func (ls *TracklistService) AddTrackToUserTrackList(userId, trackId int) error {
	err := ls.repo.AddTrackToUserTracklist(userId, trackId)
	if err != nil {
		logrus.Debug(err)
		return err
	}
	return nil
}

func (ls *TracklistService) DeleteTrackFromUserTracklist(userId, trackId int) error {
	err := ls.repo.DeleteTrackFromUserTracklist(userId, trackId)
	if err != nil {
		logrus.Debug(err)
		return err
	}
	return nil
}

func (ls *TracklistService) DeleteUserTrackList(userId int) error {
	err := ls.repo.DeleteUserTracklist(userId)
	if err != nil {
		logrus.Debug(err)
		return err
	}
	return nil
}
