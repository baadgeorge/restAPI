package service

import (
	"test/internal/datastruct"
	"test/internal/repository"
)

type IUserService interface {
	GetUser(id int) (*datastruct.User, error)
	UpdateUser(user datastruct.User) error
	DeleteUser(id int) error
	GetUserPasswordByEmail(email string) (string, error)
	GetEmailByUserId(id int) (string, error)
	GetUserIdByEmail(email string) (int, error)
}

type UserService struct {
	repo repository.IUserQuery
}

func NewUserService(repo repository.IUserQuery) IUserService {
	return &UserService{repo: repo}
}

func (us *UserService) GetUser(id int) (*datastruct.User, error) {
	user, err := us.repo.GetUser(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserService) UpdateUser(user datastruct.User) error {
	err := us.repo.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) DeleteUser(id int) error {
	err := us.repo.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) GetUserPasswordByEmail(email string) (string, error) {
	hashed_password, err := us.repo.GetUserPasswordByEmail(email)
	if err != nil {
		return "", err
	}
	return hashed_password, nil
}
func (us *UserService) GetEmailByUserId(id int) (string, error) {
	email, err := us.repo.GetEmailByUserId(id)
	if err != nil {
		return "", err
	}
	return email, nil
}
func (us *UserService) GetUserIdByEmail(email string) (int, error) {
	id, err := us.repo.GetUserIdByEmail(email)
	if err != nil {
		return 0, err
	}
	return id, nil
}
