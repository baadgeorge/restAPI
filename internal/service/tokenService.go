package service

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"time"
)

type ITokenService interface {
	NewAccessToken(userId int) (string, error)
	Parse(accessToken string) (int, error)
}

type TokenService struct {
	signingKey string
}

func NewTokenService(key string) ITokenService {
	return &TokenService{signingKey: key}
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func (ts *TokenService) NewAccessToken(userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
		},
		UserId: userId,
	})

	signedStr, err := token.SignedString([]byte(ts.signingKey))
	if err != nil {
		logrus.Debug(err)
		return "", err
	}
	return signedStr, nil

}
func (ts *TokenService) Parse(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(ts.signingKey), nil
	})
	if err != nil {
		logrus.Debug(accessToken, " ", ts.signingKey)
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		err = fmt.Errorf("can't get token claims")
		return 0, err
	}

	return claims.UserId, err
}
