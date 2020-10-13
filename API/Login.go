package API

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
)

func (a *Api) Login(user uint, pass string) (string, error) {
	storedpass, err := a.DbAccess.GetStoredPassword(user)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedpass), []byte(pass))
	if err != nil {
		return "", errors.New("Unauthenticated")
	}

	claim := Claims{
		ID: user,
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
