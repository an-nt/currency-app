package API

import (
	"CurrencyApp/Database"
	"net/http"
)

type Api struct {
	Database.Database
	apiLogin  IFmLogin
	apiExRate IFmExRate
}

type IApi interface {
	Login(w http.ResponseWriter, r *http.Request)
	GetExchangeRate(w http.ResponseWriter, r *http.Request)
}

var key = []byte("thisisasecretkey")
