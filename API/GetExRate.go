package API

import (
	"CurrencyApp/Database"

	"github.com/dgrijalva/jwt-go"
)

type ExecGetExRate struct {
	Finder Database.IGetExRate
}

type IExecGetExRate interface {
	GetExRate(token string) (uint, error)
}

func (e *ExecGetExRate) GetExRate(token string) (uint, error) {
	if !isAuthenticated(token) {
		return 0, execError("Unauthenticated")
	}
	rate, err := e.Finder.GetExRate()
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
