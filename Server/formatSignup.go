package Server

import (
	"encoding/json"
	"net/http"
)

func (sv *HttpServer) formatSignUp(w http.ResponseWriter, r *http.Request) {
	var err error
	var emp LoginInput
	var result string

	err = json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		resp := ResponseFormat{
			Message: "invalid type",
			Detail:  nil,
		}
		resp.FormatResp(w, http.StatusBadRequest)
		return
	}
	result, err = sv.Exec.SignUp(emp.ID, emp.Pass)
	if err != nil {
		resp := ResponseFormat{
			Message: err.Error(),
			Detail:  nil,
		}
		resp.FormatResp(w, http.StatusBadRequest)
		return
	}

	resp := ResponseFormat{
		Message: result,
		Detail:  nil,
	}
	resp.FormatResp(w, http.StatusOK)
	return
}
