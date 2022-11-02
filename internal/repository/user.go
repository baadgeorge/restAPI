package repository

import (
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"test/internal/datastruct"
)

type IUserQuery interface {
	CreateUser(user datastruct.FullUser) (int, error)
	GetUser(userId int) (*datastruct.User, error)
	DeleteUser(userId int) error
	UpdateUser(user datastruct.User) error
	GetUserPasswordByEmail(email string) (string, error)
	GetEmailByUserId(userId int) (string, error)
	GetUserIdByEmail(email string) (int, error)
}

type UserQuery struct {
	db *pgx.Conn
}

func NewUserQuery(db *pgx.Conn) *UserQuery {
	return &UserQuery{db: db}
}

func (u *UserQuery) CreateUser(user datastruct.FullUser) (int, error) {
	var userId int
	err := u.db.QueryRow("insert into users (first_name, last_name, email, password) values($1, $2, $3, $4) returning user_id",
		user.FirstName, user.LastName, user.Email, user.Password).Scan(&userId)
	if err != nil {
		logrus.Debug(err)
		return 0, err
	}
	return userId, nil
}
func (u *UserQuery) GetUser(userId int) (*datastruct.User, error) {
	var user datastruct.User
	err := u.db.QueryRow("select first_name, last_name, email from users where user_id=$1",
		userId).Scan(&user.FirstName, &user.LastName, &user.Email)
	user.ID = userId
	if err != nil {
		logrus.Debug(err)
		return nil, err
	}
	return &user, nil
}

func (u *UserQuery) DeleteUser(userId int) error {
	_, err := u.db.Exec("delete from users where user_id=$1", userId)
	if err != nil {
		logrus.Debug(err)
		return err
	}
	return nil
}

func (u *UserQuery) UpdateUser(user datastruct.User) error {
	_, err := u.db.Exec("update users set first_name=$1, last_name=$2, email=$3 where user_id=$4",
		user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		logrus.Debug(err)
		return err
	}
	return nil
}

func (u *UserQuery) GetUserPasswordByEmail(email string) (string, error) {
	var password string
	err := u.db.QueryRow("select password from users where email=$1", email).Scan(&password)
	if err != nil {
		logrus.Debug(err)
		return "", err
	}
	return password, nil
}

func (u *UserQuery) GetEmailByUserId(userId int) (string, error) {
	var email string
	err := u.db.QueryRow("select email from users where user_id=$1", userId).Scan(&email)
	if err != nil {
		logrus.Debug(err)
		return "", err
	}
	return email, nil
}

func (u *UserQuery) GetUserIdByEmail(email string) (int, error) {
	var userId int
	err := u.db.QueryRow("select user_id from users where email=$1", email).Scan(&userId)
	if err != nil {
		logrus.Debug(err)
		return 0, err
	}
	return userId, nil
}
