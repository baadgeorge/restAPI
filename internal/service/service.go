package service

import (
	"test/internal/repository"
)

type DataService struct {
	IAuthService
	ITrackService
	IUserService
	ITracklistService
	ITokenService
}

func NewService(repo *repository.Repository, tokenServ ITokenService) *DataService {
	return &DataService{
		IAuthService:      NewAuthService(repo.IUserQuery, tokenServ),
		ITrackService:     NewTrackService(repo.ITrackQuery),
		IUserService:      NewUserService(repo.IUserQuery),
		ITracklistService: NewTrackListService(repo.ITracklistQuery),
		ITokenService:     tokenServ,
	}
}
