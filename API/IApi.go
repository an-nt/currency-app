package API

import (
	"CurrencyApp/Database"
)

type Api struct {
	DbAccess Database.IDatabaseAccess
}

type IApi interface {
	Login(user uint, pass string) (string, error)
	GetExRate(token string) (uint, error)
}
