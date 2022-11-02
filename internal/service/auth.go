package service

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"test/internal/datastruct"
	"test/internal/repository"
)

type IAuthService interface {
	SignUp(user datastruct.FullUser) (int, error)
	Login(email, reqPassword string) (string, error)
}

type AuthService struct {
	repo         repository.IUserQuery
	tokenService ITokenService
}

func NewAuthService(repo repository.IUserQuery, tokenService ITokenService) IAuthService {
	return &AuthService{repo: repo, tokenService: tokenService}
}

func (a *AuthService) SignUp(user datastruct.FullUser) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	user.Password = string(hashedPassword)
	userId, err := a.repo.CreateUser(user)
	if err != nil {
		return 0, err
	}
	return userId, nil

}
func (a *AuthService) Login(email, reqPassword string) (string, error) {
	password, err := a.repo.GetUserPasswordByEmail(email)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(reqPassword))
	if err != nil {
		logrus.Debug(password, reqPassword)
		return "", err
	}
	userId, err := a.repo.GetUserIdByEmail(email)
	if err != nil {
		return "", err
	}
	logrus.Print(a.tokenService.NewAccessToken(userId))
	jwt, err := a.tokenService.NewAccessToken(userId)
	if err != nil {
		return "", err
	}
	return jwt, nil

}
