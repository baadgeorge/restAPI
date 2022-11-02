package service

import (
	"test/internal/datastruct"
	"test/internal/repository"
)

type ITrackService interface {
	GetTrackById(id int) (*datastruct.Track, error)
	GetTracksByAuthorAndTitle(author string, title string) ([]datastruct.Track, error)
	GetTracksByAuthor(author string) ([]datastruct.Track, error)
	GetTracksByTitle(title string) ([]datastruct.Track, error)
	UpdateTrack(track datastruct.Track) error
	DeleteTrack(id int) error
	CreateTrack(track *datastruct.Track) (int, error)
}

type TrackService struct {
	repo repository.ITrackQuery
}

func NewTrackService(repo repository.ITrackQuery) ITrackService {
	return &TrackService{repo: repo}
}

func (ts *TrackService) GetTrackById(id int) (*datastruct.Track, error) {
	track, err := ts.repo.GetTrackById(id)
	if err != nil {
		return nil, err
	}
	return track, nil
}

func (ts *TrackService) GetTracksByAuthorAndTitle(author string, title string) ([]datastruct.Track, error) {
	tracks, err := ts.repo.GetTracksByAuthorAndTitle(author, title)
	if err != nil {
		return nil, err
	}
	return tracks, nil
}

func (ts *TrackService) GetTracksByAuthor(author string) ([]datastruct.Track, error) {
	tracks, err := ts.repo.GetTracksByAuthor(author)
	if err != nil {
		return nil, err
	}
	return tracks, nil
}
func (ts *TrackService) GetTracksByTitle(title string) ([]datastruct.Track, error) {
	tracks, err := ts.repo.GetTracksByTitle(title)
	if err != nil {
		return nil, err
	}
	return tracks, nil
}

func (ts *TrackService) UpdateTrack(track datastruct.Track) error {
	err := ts.repo.UpdateTrack(track)
	if err != nil {
		return err
	}
	return nil
}
func (ts *TrackService) DeleteTrack(id int) error {
	err := ts.repo.DeleteTrack(id)
	if err != nil {
		return err
	}
	return nil
}

func (ts *TrackService) CreateTrack(track *datastruct.Track) (int, error) {
	trackId, err := ts.repo.CreateTrack(*track)
	if err != nil {
		return 0, err
	}
	return trackId, nil
}
