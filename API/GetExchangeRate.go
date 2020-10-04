package API

import (
	"CurrencyApp/Database"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func (a *Api) GetExchangeRate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "JSON")

	if a.apiExRate == nil {
		a.apiExRate = &FmExRate{
			db: a.Db,
		}
	}

	resp, err := a.apiExRate.FmExRate(r)
	if err != nil {
		resp = FmResp{
			Statuscode: http.StatusBadRequest,
			Message:    err.Error(),
			Detail:     nil,
		}
		w.WriteHeader(resp.Statuscode)
		json.NewEncoder(w).Encode(resp)
		return
	}

	w.WriteHeader(resp.Statuscode)
	json.NewEncoder(w).Encode(resp)
}

type FmExRate struct {
	IAuth
	db *sql.DB
}
type IFmExRate interface {
	FmExRate(r *http.Request) (FmResp, error)
}

func (f *FmExRate) FmExRate(r *http.Request) (FmResp, error) {
	var resp FmResp
	token := r.Header.Get("Authorization")

	if f.IAuth == nil {
		f.IAuth = &Auth{
			db: f.db,
		}
	}

	rate, valid := f.IAuth.Auth(token)
	if !valid {
		resp = FmResp{
			Statuscode: http.StatusForbidden,
			Message:    "Unauthenticated",
			Detail:     nil,
		}
		return resp, nil
	}
	resp = FmResp{
		Statuscode: http.StatusOK,
		Message:    "Success",
		Detail:     rate,
	}
	return resp, nil
}

type Auth struct {
	Database.IGetExRate
	db *sql.DB
}
type IAuth interface {
	Auth(token string) (uint, bool)
}

func (a *Auth) Auth(token string) (uint, bool) {
	if token == "" {
		return 0, false
	}

	nullclaim := &Claims{}
	_, err := jwt.ParseWithClaims(token, nullclaim, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return 0, false
	}

	if a.IGetExRate == nil {
		a.IGetExRate = &Database.MSSQL{
			DB: Database.DB{
				Db: a.db,
			},
		}
	}
	rate, err := a.GetExRate()
	if err != nil {
		return 0, false
	}
	return rate, true
}
