package API

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

func (a *Api) GetExRate(token string) (uint, error) {
	if !isAuthenticated(token) {
		return 0, errors.New("Unauthenticated")
	}
	rate, err := a.DbAccess.GetExRate()
	if err != nil {
		return 0, err
	}
	return rate, nil
}

func isAuthenticated(token string) bool {
	claim := &Claims{}
	_, err := jwt.ParseWithClaims(token, claim, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err == nil {
		return true
	}
	return false
}
