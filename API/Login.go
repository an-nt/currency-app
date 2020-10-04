package API

import (
	"CurrencyApp/Database"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type ExecLogin struct {
	Checker Database.ICheckExist
}
type IExecLogin interface {
	Login(user uint, pass string) (string, error)
}

func (e *ExecLogin) Login(user uint, pass string) (string, error) {
	staff, err := e.Checker.CheckExist(user, pass)
	if err != nil {
		return "", err
	}
	claim := Claims{
		ID:          staff.ID,
		Name:        staff.FullName,
		Male:        staff.Male,
		Nationality: staff.Nationality,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(300 * time.Second).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(key)
}

type Claims struct {
	ID          uint
	Name        string
	Male        bool
	Nationality string
	jwt.StandardClaims
}

var key = []byte("thisisasecretkey")
