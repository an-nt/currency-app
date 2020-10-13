package Server

import (
	"encoding/json"
	"net/http"
)

type LoginInput struct {
	ID   uint   `json:"id"`
	Pass string `json:"pass"`
}

func (sv *HttpServer) FormatLogin(w http.ResponseWriter, r *http.Request) {
	var input LoginInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		resp := ResponseFormat{
			Message: "invalid type",
			Detail:  nil,
		}
		resp.FormatResp(w, http.StatusBadRequest)
		return
	}

	token, err := sv.Exec.Login(input.ID, input.Pass)
	if err != nil {
		resp := ResponseFormat{
			Message: err.Error(),
			Detail:  nil,
		}
		resp.FormatResp(w, http.StatusUnauthorized)
		return
	}

	resp := ResponseFormat{
		Message: "Success",
		Detail:  token,
	}
	resp.FormatResp(w, http.StatusOK)
}
