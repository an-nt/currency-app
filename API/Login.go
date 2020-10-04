package API

import (
	"CurrencyApp/Database"
	"CurrencyApp/Model"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func (a *Api) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "JSON")

	if a.apiLogin == nil {
		a.apiLogin = &FmLogin{
			db: a.DB,
		}
	}

	resp, err := a.apiLogin.FmLogin(r)
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

type FmLogin struct {
	db Database.DB
	IToken
}
type IFmLogin interface {
	FmLogin(r *http.Request) (FmResp, error)
}

func (f *FmLogin) FmLogin(r *http.Request) (FmResp, error) {
	var staff Model.Employee
	var resp FmResp

	err := json.NewDecoder(r.Body).Decode(&staff)
	if err != nil {
		return resp, err
	}

	if f.IToken == nil {
		f.IToken = &Token{
			db: f.db,
		}
	}
	token, err := f.GenerateToken(staff)
	if err != nil {
		resp = FmResp{
			Statuscode: http.StatusForbidden,
			Message:    err.Error(),
			Detail:     nil,
		}
		return resp, nil
	}

	resp = FmResp{
		Statuscode: http.StatusOK,
		Message:    "Authorized",
		Detail:     token,
	}
	return resp, nil
}

type Token struct {
	db      Database.DB
	checker Database.ICheck //actual checker
}
type IToken interface {
	GenerateToken(staff Model.Employee) (string, error)
}

func (t *Token) GenerateToken(staff Model.Employee) (string, error) {
	if t.checker == nil {
		t.checker = &Database.MSSQL{
			DB: t.db,
		}
	}
	employee, err := t.checker.Check(staff)
	if err != nil || employee.ID == 0 || employee.FullName == "" {
		return "", err
	}

	claim := Claims{
		ID:   employee.ID,
		Name: employee.FullName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(300 * time.Second).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(key)
}
