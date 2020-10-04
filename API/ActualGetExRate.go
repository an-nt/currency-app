package API

import (
	"CurrencyApp/Database"

	"github.com/dgrijalva/jwt-go"
)

type ExecGetExRate struct {
	finder Database.IGetExRate
}

type IExecGetExRate interface {
	GetExRate(token string) (uint, error)
}

func (e *ExecGetExRate) GetExRate(token string) (uint, error) {
	if !isAuthenticated(token) {
		return 0, execError("Unauthenticated")
	}
	rate, err := e.finder.GetExRate()
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
	if err != nil {
		return false
	} else {
		return true
	}
}

func execError(description string) error {
	return &problem{
		detail: description,
	}
}

type problem struct {
	detail string
}

func (e *problem) Error() string {
	return e.detail
}
