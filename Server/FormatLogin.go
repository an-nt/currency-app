package Server

import (
	"CurrencyApp/API"
	"encoding/json"
	"net/http"
)

type LoginInput struct {
	ID   uint   `json:"id"`
	Pass string `json:"pass"`
}
type LoginFormatter struct {
	Handler API.IExecLogin
}

type ILoginFormatter interface {
	FormatLogin(w http.ResponseWriter, r *http.Request)
}

func (f *LoginFormatter) FormatLogin(w http.ResponseWriter, r *http.Request) {
	var input LoginInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		resp := FmResp{
			Message: "invalid type",
			Detail:  nil,
		}
		resp.FormatResp(w, http.StatusBadRequest)
		return
	}

	token, err := f.Handler.Login(input.ID, input.Pass)
	if err != nil {
		resp := FmResp{
			Message: err.Error(),
			Detail:  nil,
		}
		resp.FormatResp(w, http.StatusUnauthorized)
		return
	}

	resp := FmResp{
		Message: "Success",
		Detail:  token,
	}
	resp.FormatResp(w, http.StatusOK)
}
